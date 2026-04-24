package ledger

import (
	"DIstributed-Ledger-Raft1/internal/ledger/model"
	"errors"
	"sync"
)

type AccountManager struct {
	Mu      sync.RWMutex
	Account map[string]*model.Account //key 公钥地址
}

func NewAccountManager() *AccountManager {
	return &AccountManager{
		Account: make(map[string]*model.Account),
	}
}

// 开户 金额默认0
func (am *AccountManager) CreateAccount(address string) error {
	am.Mu.Lock()
	defer am.Mu.Unlock()

	if _, exists := am.Account[address]; exists {
		return errors.New("账户已存在")
	}
	am.Account[address] = &model.Account{
		Address:  address,
		Balance:  0,
		IsFrozen: false,
	}
	return nil
}

// 查询金额
func (am *AccountManager) GetBalance(address string) (int64, error) {
	am.Mu.RLock()
	defer am.Mu.RUnlock()

	acc, exists := am.Account[address]
	if !exists {
		return 0, errors.New("账户不存在")
	}
	return acc.Balance, nil
}

// 冻结风险账户
func (am *AccountManager) FreezeAccount(address string) error {
	am.Mu.RLock()
	defer am.Mu.RUnlock()
	acc, exists := am.Account[address]
	if !exists {
		return errors.New("该账户不存在")
	}
	acc.IsFrozen = true
	return nil
}

// 初始资产配置 用于测试
func (am *AccountManager) InjectAssets(address string, account int64) error {
	am.Mu.Lock()
	defer am.Mu.Unlock()

	acc, exists := am.Account[address]
	if !exists {
		return errors.New("该账户不存在")
	}
	acc.Balance += account
	return nil
}
