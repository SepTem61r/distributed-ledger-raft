package model

type Account struct {
	Address  string //账户地址
	Balance  int64  //账户余额
	IsFrozen bool   //风险评估 true不能交易
}
