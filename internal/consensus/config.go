package consensus

import (
	"time"

	"github.com/hashicorp/raft"
)

func NewConfig(localID string) *raft.Config {

	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localID)
	config.HeartbeatTimeout = 1000 * time.Millisecond
	config.ElectionTimeout = 1000 * time.Millisecond
	config.CommitTimeout = 500 * time.Millisecond
	return config
}
