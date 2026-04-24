# 第一阶段：构建 Go 二进制文件
FROM golang:1.26-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置 Go 代理，加速国内依赖下载
ENV GOPROXY=https://goproxy.cn,direct

# 先复制 go.mod 和 go.sum 下载依赖（利用 Docker 缓存机制）
COPY go.mod go.sum ./
RUN go mod download

# 复制整个项目源码并编译 ledger-node
COPY . .
# 编译出可执行文件 ledger-node
RUN go build -o ledger-node ./cmd/ledger-node/main.go


# 第二阶段：运行环境（使用极简的 alpine 镜像减小体积）
FROM alpine:latest

WORKDIR /app

# 从 builder 阶段把编译好的程序拷过来
COPY --from=builder /app/ledger-node .

# 创建数据存储目录，Raft 会把日志写在这里
RUN mkdir -p /app/data

# 暴露可能用到的端口（声明作用，实际映射由 docker-compose 决定）
EXPOSE 8081 8082 8083 7001 7002 7003

# 容器启动的默认命令，具体参数会被 docker-compose 的 command 覆盖
ENTRYPOINT ["./ledger-node"]