package consensus

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/hashicorp/raft"
)

func NewTransport(bindAddr string) (raft.Transport, error) {
	addr, err := net.ResolveTCPAddr("tcp", bindAddr)
	if err != nil {
		return nil, fmt.Errorf("解析TCP地址失败: %v", err)
	}
	transport, err := raft.NewTCPTransport(bindAddr, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("创建TCP传输层失败: %v", err)
	}
	return transport, nil
}
