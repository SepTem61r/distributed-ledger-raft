# 🔗 Distributed-Ledger-Raft

> 基于 Raft 强一致性共识构建的高性能分布式审计账本系统。

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)
![Consensus](https://img.shields.io/badge/Consensus-HashiCorp_Raft-blue?style=flat-square)
![Docker](https://img.shields.io/badge/Docker-Supported-2496ED?style=flat-square&logo=docker)
![Status](https://img.shields.io/badge/Status-Active-success?style=flat-square)

本项目存储层采用级联哈希存证技术，实现记录链路可追溯、不可篡改；网络层与状态机通过 `HashiCorp Raft` 算法保证多节点强一致性；并通过 `ECDSA` 异步签名确保指令权属与数据完整。

在多节点环境下，系统兼顾线性一致性与金融级高可用，配合完善的权限治理机制，实现对账内资产全生命周期的合规化、透明化管理。

---

## 核心特性

- **强一致性转账**：基于 Raft 算法，确保多节点状态机（FSM）数据严格一致收敛。
- **不可篡改与可追溯**：级联哈希链式结构，保障金融级数据防伪造安全。
- **高可用与故障自愈**：集群支持节点宕机重连，自动通过 Snapshot 快照与日志回放恢复最新状态。
- **资产与风控治理**：支持初始资产配置、原子化转账、ECDSA 身份验签及风险账户实时冻结拦截。

---

## 快速开始

### 方式一：Docker Compose 一键启动（推荐）

要求：已安装 `Docker` 与 `Docker Compose`。

```bash
# 1. 克隆代码库
git clone [https://github.com/SepTem61r/distributed-ledger-raft.git]
cd distributed-ledger-raft

# 2. 一键拉起 3 节点分布式集群
docker compose up -d --build

# 3. 运行本地钱包测试脚本，验证强一致转账
go run ./cmd/wallet/main.go
```
### 方式二：本地源码启动
```bash
# 开启三个终端，分别启动 3 个集群节点
go run ./cmd/ledger-node/main.go -id=node1 -http=8081 -raft=7001
go run ./cmd/ledger-node/main.go -id=node2 -http=8082 -raft=7002
go run ./cmd/ledger-node/main.go -id=node3 -http=8083 -raft=7003

# 开启第四个终端，执行钱包客户端流转
go run ./cmd/wallet/main.go
```
## 项目结构
```text
distributed-ledger-raft/
├── cmd/                          # 程序入口
│   ├── ledger-node/              # 服务端：组装依赖，启动 API 和 Raft 节点
│   └── wallet/                   # 客户端：模拟央行空投、用户开户与转账测试
├── internal/                     # 核心业务逻辑
│   ├── consensus/                # [Raft 共识层] 管理集群状态机(FSM)与快照(Snapshot)
│   ├── ledger/                   # [账本核心层] 核心转账逻辑、账户管理、区块与交易模型
│   ├── crypto/                   # [密码学安全] ECDSA 签名验签、SHA256 哈希计算
│   └── storage/                  # [存储层] Raft 本地日志数据库对接
├── api/                          #网络接口层
│   └── http/
│       ├── dto/                  # 统一请求与响应数据模型 (Request/Response)
│       ├── handler/              # 路由控制器 (处理参数验证与底层路由)
│       └── router.go             # Gin 框架路由注册
├── data/                         # 运行数据卷 (BoltDB 数据库与快照，由 Docker 挂载)
├── Dockerfile                    # 节点容器构建配置
├── docker-compose.yml            # 3 节点集群编排配置
└── go.mod                        # Go 依赖管理
```
