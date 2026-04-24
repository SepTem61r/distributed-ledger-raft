package consensus

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

func SetupRaftNode(localID string, bindAddr string, dataDir string, fsm raft.FSM) (*raft.Raft, error) {

	config := NewConfig(localID)

	transport, err := NewTransport(bindAddr)
	if err != nil {
		return nil, fmt.Errorf("创建网络通信信道失败%v", err)
	}

	snapshort, err := raft.NewFileSnapshotStore(dataDir, 2, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("创建快照存储失败: %v", err)
	}

	boltDB, err := raftboltdb.NewBoltStore(filepath.Join(dataDir, "raft.db"))
	if err != nil {
		return nil, fmt.Errorf("创建BoltDB存储失败: %v", err)
	}
	ra, err := raft.NewRaft(config, fsm, boltDB, boltDB, snapshort, transport)
	if err != nil {
		return nil, fmt.Errorf("启动Raft节点失败: %v", err)
	}
	fmt.Printf("Raft 节点 [%s] 启动成功，内部通信端口: %s\n", localID, bindAddr)
	return ra, nil
}
