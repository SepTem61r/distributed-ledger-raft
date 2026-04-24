package handler

import (
	"DIstributed-Ledger-Raft1/internal/crypto"
	"DIstributed-Ledger-Raft1/internal/ledger/model"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/raft"

	"DIstributed-Ledger-Raft1/api/http/dto"
	"DIstributed-Ledger-Raft1/internal/ledger"
)

// 需要 Raft 节点来写数据，需要 AccountManager 来读数据
type LedgerHandler struct {
	RaftNode *raft.Raft
	Account  *ledger.AccountManager
}

// 查询节点状态
func (h *LedgerHandler) GetStatus(c *gin.Context) {
	dto.Success(c, gin.H{
		"state":  h.RaftNode.State().String(),
		"leader": h.RaftNode.Leader(),
	})
}

// 查询余额
func (h *LedgerHandler) GetBalance(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		dto.Error(c, 400, "address不为空")
		return
	}
	bal, err := h.Account.GetBalance(addr)
	if err != nil {
		dto.Error(c, 404, "账户不存在")
		return
	}
	dto.Success(c, gin.H{"address": addr[:15] + "...", "balance": bal})
}

// 发起转账 (必须扔进 Raft 共识)
func (h *LedgerHandler) Transfer(c *gin.Context) {
	// 1. 只有 Leader 才能接收写请求
	if h.RaftNode.State() != raft.Leader {
		dto.Error(c, 403, "拒绝写入：请将请求发送给 Leader 节点")
		return
	}

	var req *dto.TransferReq
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.Error(c, 400, "参数错误:"+err.Error())
		return
	}
	sig, _ := hex.DecodeString(req.Signature)
	// 将 DTO 转换为底层 Model
	tx := model.Transaction{
		From:      req.From,
		To:        req.To,
		Amount:    req.Amount,
		Timestamp: time.Unix(0, req.Timestamp),
		Signature: sig,
	}
	tx.TxID = crypto.HashData(tx.From, tx.To, tx.Amount, tx.Timestamp.UnixNano())
	// 打包投递给 Raft 引擎
	txBytes, _ := json.Marshal(tx)
	future := h.RaftNode.Apply(txBytes, 5*time.Second)
	// 查看业务状态机是否拒绝
	if err := future.Error(); err != nil {
		dto.Error(c, 500, "Raft 共识失败: "+err.Error())
		return
	}

	if bizErr := future.Response(); bizErr != nil {
		dto.Error(c, 400, fmt.Sprintf("业务规则拒绝: %v", bizErr))
		return
	}
	dto.Success(c, gin.H{"tx_id": "已上链"})
}
