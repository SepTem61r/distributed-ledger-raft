package http

import (
	"DIstributed-Ledger-Raft1/api/http/handler"
	"DIstributed-Ledger-Raft1/internal/ledger"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/raft"
)

func StartServer(port string, raftNode *raft.Raft, accountManager *ledger.AccountManager) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	h := &handler.LedgerHandler{
		RaftNode: raftNode,
		Account:  accountManager,
	}
	api := r.Group("/api/v1")
	{
		api.GET("/status", h.GetStatus)
		api.GET("/balance", h.GetBalance)
		api.POST("/transfer", h.Transfer)
	}
	fmt.Printf("Gin API 服务已启动，监听端口: %s\n", port)
	return r.Run(":" + port)
}
