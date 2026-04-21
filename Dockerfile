# 构建阶段
FROM golang:1.22-alpine AS builder

# 安装必要的构建依赖
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o web3-ai-bsc ./cmd/main.go

# 运行阶段
FROM alpine:latest

# 安装运行时依赖
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/web3-ai-bsc .

# 复制环境变量文件（可选，实际部署时建议通过docker-compose或k8s配置）
COPY --from=builder /app/env.example .env.example

# 暴露端口
EXPOSE 8080

# 设置非root用户运行
RUN adduser -D -g '' appuser
USER appuser

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 启动应用
CMD ["./web3-ai-bsc"]