package dto

// 接收外部转账请求的格式
type TransferReq struct {
	From      string `json:"from" binding:"required"`
	To        string `json:"to" binding:"required"`
	Amount    int64  `json:"amount" binding:"required,gt=0"`
	Timestamp int64  `json:"timestamp" binding:"required"`
	Signature string `json:"signature" binding:"required"` // hex格式的签名
}
