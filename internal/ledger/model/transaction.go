package model

import "time"

type Transaction struct {
	TxID      string    //交易id
	From      string    //付款方
	To        string    //收款方
	Amount    int64     //交易金额
	Timestamp time.Time //交易时间
	Signature []byte    //ecdsa 签名
}
