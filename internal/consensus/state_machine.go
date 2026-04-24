package consensus

import (
	"DIstributed-Ledger-Raft1/internal/ledger"
	"DIstributed-Ledger-Raft1/internal/ledger/model"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/raft"
)

type LedgerFSM struct {
	AccountMAnager *ledger.AccountManager
}
type LedgerSnapshot struct {
	snapshotData []byte
}

func NewLedgerFSM(am *ledger.AccountManager) *LedgerFSM {
	return &LedgerFSM{
		AccountMAnager: am,
	}
}
func (fsm *LedgerFSM) Apply(log *raft.Log) interface{} {
	var tx model.Transaction
	if err := json.Unmarshal(log.Data, &tx); err != nil {
		return fmt.Errorf("解析日志失败: %v", err)
	}
	if tx.From == "SYSTEM" {
		_ = fsm.AccountMAnager.CreateAccount(tx.To)
		_ = fsm.AccountMAnager.InjectAssets(tx.To, tx.Amount)
		fmt.Printf("[系统空投] 为 %s 注入 %d 资产已同步\n", tx.To[:10], tx.Amount)
		return nil
	}
	err := fsm.AccountMAnager.ExecuteTransfer(tx)
	if err != nil {
		fmt.Printf("[FSM Apply 失败] TxID: %s, 原因: %v\n", tx.TxID, err)
		return err
	}

	return nil
}
func (fsm *LedgerFSM) Snapshot() (raft.FSMSnapshot, error) {
	fsm.AccountMAnager.Mu.RLock()
	defer fsm.AccountMAnager.Mu.RUnlock()
	data, err := json.Marshal(fsm.AccountMAnager.Account)
	if err != nil {
		return nil, err
	}

	return &LedgerSnapshot{snapshotData: data}, nil
}
func (s *LedgerSnapshot) Persist(sink raft.SnapshotSink) error {
	_, err := sink.Write(s.snapshotData)
	if err != nil {
		sink.Cancel()
		return err
	}
	if err := sink.Close(); err != nil {
		return fmt.Errorf("快照刷盘失败: %v", err)
	}
	return nil
}

func (s *LedgerSnapshot) Release() {}
func (fsm *LedgerFSM) Restore(rc io.ReadCloser) error {
	defer rc.Close()
	data, err := io.ReadAll(rc)
	if err != nil {
		return err
	}
	var restoredAccounts map[string]*model.Account
	if err := json.Unmarshal(data, &restoredAccounts); err != nil {
		return err
	}
	fsm.AccountMAnager.Mu.Lock()
	defer fsm.AccountMAnager.Mu.Unlock()
	fsm.AccountMAnager.Account = restoredAccounts
	fmt.Println("[FSM Restore] 已从快照成功恢复账本状态")
	return nil
}
