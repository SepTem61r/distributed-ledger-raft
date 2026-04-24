基于 Raft 算法的分布式账本技术栈：Go+Hash 链 +Raft 共识 + Docker
项目简介：本项目基于 Raft 强一致性共识构建高性能分布式审计系统。存储层采用级联哈希存证技术，实现记录链路可追溯、不可篡改。通过 ECDSA 异步签名确保指令权属与数据完整。
系统支持初始资产配置、原子化转账、实时余额对账及风控准入冻结等核心功能。
在多节点环境下，系统兼顾线性一致性与金融级高可用，配合完善的权限治理机制，实现对账内资产全生命周期的合规化、透明化管理。
亮点：强一致转账、不可篡改与可追溯、高可用与故障自愈、一致性收敛能力

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
