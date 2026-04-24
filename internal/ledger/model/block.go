package model

import (
	"DIstributed-Ledger-Raft1/internal/crypto"
	"time"
)

type Block struct {
	Index       uint64        //区块高度
	Timestamp   time.Time     //出块时间
	Transaction []Transaction // 此区块包含的交易记录
	PrevHash    string        //上一区块的哈希
	Hash        string        // 当前区块的哈希
}

func (b *Block) GenerateHash() string {
	return crypto.HashData(
		b.Index,
		b.Timestamp.UnixNano(),
		b.Transaction,
		b.PrevHash)
}
