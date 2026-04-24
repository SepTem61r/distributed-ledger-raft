distributed-ledger-raft
├── cmd/                          # 程序入口
│   └── ledger-node/              
│   │   └── main.go               # 唯一入口：初始化配置、组装依赖、启动 API 和 Raft 节点
│   └──wallet/
│   │  └──main.go                 # 测试转账
├── internal/                     # 核心业务逻辑
│   ├── consensus/                # Raft 共识层 
│   │   ├── raft_node.go          # 初始化 hashicorp/raft，管理集群节点状态
│   │   ├── state_machine.go      # 实现 raft.FSM 接口：Apply(执行交易)、Snapshot(快照)
│   │   ├── transport.go          # Raft 节点间的 TCP 通信层实现
│   │   └── config.go             # Raft 共识算法库的核心配置初始化逻辑
│   ├── ledger/                   # 账本核心层 (直接承接业务)
│   │   ├── model/                
│   │   │   ├── block.go          # 定义账户模型
│   │   │   ├── transaction.go    # 定义交易模型
│   │   │   └── account.go        # 定义区块模型
│   │   ├── chain.go             
│   │   ├── account.go            # 账户开户、状态变更
│   │   └── transfer.go           # 核心转账逻辑 (扣款、加钱)
│   ├── crypto/                  
│   │   ├── ecdsa.go              # 签名与验签
│   │   └── hash.go               # SHA256 与前置哈希计算
│   └── api/                      # 网络接口层
│       └── http/                 
│           ├── dto/             
│           │   ├── request.go    # 接收外部参数 (TransferReq)，带 binding 校验
│           │   └── response.go   # 返回给外部的格式化JSON结构
│           ├── handler/          # 路由处理器 (负责解析 DTO -> 调用 ledger -> 返回 DTO)
│           │   └── tx_handler.go
│           └── router.go         # Gin 路由注册
├── Dockerfile                    # 容器化部署
├── docker-compose.yml            # 一键拉起 3 节点集群
├── go.mod
└── README.md
