package ledger

import (
	"DIstributed-Ledger-Raft1/internal/crypto"
	"DIstributed-Ledger-Raft1/internal/ledger/model"
	"errors"
)

// 执行转账交易
func (am *AccountManager) ExecuteTransfer(tx model.Transaction) error {

	if !crypto.Verity(tx.From, tx.TxID, tx.Signature) {
		return errors.New("签名校验失败：非法的转账请求")
	}

	am.Mu.RLock()
	defer am.Mu.RUnlock()

	fromAcc, exists := am.Account[tx.From]
	if !exists {
		return errors.New("付款账户不存在")
	}
	toAcc, exists := am.Account[tx.To]
	if !exists {
		return errors.New("收款账户不存在")
	}
	if fromAcc.IsFrozen == true {
		return errors.New("付款用户已被冻结")
	}
	if toAcc.IsFrozen == true {
		return errors.New("收款用户已被冻结")
	}
	if tx.Amount <= 0 {
		return errors.New("转账金额必须大于0")
	}
	if fromAcc.Balance < tx.Amount {
		return errors.New("余额不足")
	}

	fromAcc.Balance -= tx.Amount
	toAcc.Balance += tx.Amount
	return nil
}
