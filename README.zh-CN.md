# Web3 AI BSC Agent

<div align="center">

[![English](https://img.shields.io/badge/English-README-blue)](README.md)
[![中文](https://img.shields.io/badge/中文-README-red)](README.zh-CN.md)

![Go](https://img.shields.io/badge/Go-1.22+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Docker](https://img.shields.io/badge/Docker-Supported-orange.svg)
![Web3](https://img.shields.io/badge/Web3-Ethereum-purple.svg)

**一个生产级的 Web3 AI 代理服务，支持 BSC 链上事件监听、ERC20 合约自动解析、AI 风险决策和自动链上操作**

</div>

## 📋 目录

- [项目简介](#-项目简介)
- [技术栈](#-技术栈)
- [核心功能](#-核心功能)
- [快速开始](#-快速开始)
- [API 文档](#-api-文档)
- [部署指南](#-部署指南)
- [项目结构](#-项目结构)
- [贡献指南](#-贡献指南)
- [许可证](#-许可证)

## 🎯 项目简介

Web3 AI BSC Agent 是一个创新的区块链智能代理系统，结合了 Go 语言的高性能、BSC 链的低成本优势以及人工智能的决策能力。该系统能够实时监控链上活动，自动解析 ERC20 代币信息，并通过 AI 模型对交易进行风险评估和决策。

### 应用场景
- 🔍 **风险监控**: 实时检测可疑交易模式
- 🤖 **自动化交易**: 基于 AI 决策执行链上操作
- 📊 **数据分析**: 收集和分析链上数据
- 🛡️ **安全防护**: 识别潜在的欺诈行为

## 💻 技术栈

| 类别 | 技术 |
|------|------|
| **编程语言** | Go 1.22+ |
| **Web 框架** | Gin Framework |
| **区块链交互** | go-ethereum |
| **AI 集成** | OpenAI API / 兼容的 LLMs |
| **容器化** | Docker |
| **配置管理** | 环境变量 |

## ✨ 核心功能

- ✅ **BSC 链连接**: 稳定可靠的 BSC 节点连接
- ✅ **ERC20 自动解析**: 自动获取代币名称、符号、精度等信息
- ✅ **实时事件监听**: 监听 Transfer 事件并实时响应
- ✅ **AI 风险决策**: 基于大语言模型的交易风险评估
- ✅ **RESTful API**: 提供完整的 HTTP API 接口
- ✅ **Docker 部署**: 一键容器化部署
- ✅ **错误处理**: 完善的错误处理和日志记录
- ✅ **优雅关闭**: 支持平滑重启和关闭

## 🚀 快速开始

### 前置要求

- Go 1.22+
- Docker (可选)
- BSC RPC 节点访问权限
- OpenAI API Key 或兼容的 LLM API

### 本地运行

1. **克隆项目**
   ```bash
   git clone https://github.com/JackBean2134/web3-ai-bsc-agent.git
   cd web3-ai-bsc-agent
   ```

2. **配置环境变量**
   ```bash
   cp env.example .env
   # 编辑 .env 文件，填入你的配置
   ```

3. **安装依赖**
   ```bash
   go mod download
   ```

4. **运行服务**
   ```bash
   go run cmd/main.go
   ```

5. **访问 API**
   ```
   http://localhost:8080/health
   ```

### Docker 运行

```bash
# 构建镜像
docker build -t web3-ai-bsc-agent .

# 运行容器
docker run -d \
  --name web3-ai-bsc-agent \
  -p 8080:8080 \
  --env-file .env \
  web3-ai-bsc-agent
```

## 📡 API 文档

### 健康检查

```bash
GET /health
```

**响应示例:**
```json
{
  "status": "ok",
  "message": "Web3 AI BSC Agent is running"
}
```

### 查询 BNB 余额

```bash
GET /balance/:address
```

**参数:**
- `address`: Ethereum 地址

**响应示例:**
```json
{
  "address": "0x1234567890abcdef1234567890abcdef12345678",
  "bnb_wei": "1000000000000000000"
}
```

### 查询 ERC20 代币信息

```bash
GET /erc20/info?contract=0x...
```

**参数:**
- `contract`: ERC20 合约地址

**响应示例:**
```json
{
  "contract": "0x...",
  "name": "Binance USD",
  "symbol": "BUSD",
  "decimals": 18
}
```

### 查询 ERC20 代币余额

```bash
GET /erc20/balance?contract=0x...&address=0x...
```

**参数:**
- `contract`: ERC20 合约地址
- `address`: 钱包地址

**响应示例:**
```json
{
  "contract": "0x...",
  "address": "0x...",
  "balance": "1000000000000000000"
}
```

### AI 交易决策

```bash
GET /ai/decision?from=0x...&to=0x...&amount=100
```

**参数:**
- `from`: 发送方地址
- `to`: 接收方地址
- `amount`: 交易金额（可选）

**响应示例:**
```json
{
  "from": "0x...",
  "to": "0x...",
  "amount": "100",
  "decision": {
    "approve": true,
    "risk_level": "low",
    "reason": "Normal transaction pattern"
  }
}
```

## 🐳 部署指南

### 环境变量配置

创建 `.env` 文件并配置以下变量：

```env
# BSC 节点配置
BSC_RPC_URL=https://bsc-dataseed.binance.org/

# AI 配置
OPENAI_API_KEY=your-api-key
OPENAI_MODEL=gpt-3.5-turbo

# 服务配置
SERVER_PORT=8080
LOG_LEVEL=info

# 监听的合约地址（可选）
WATCH_CONTRACT=0x...
```

### 生产环境部署

1. **使用 Docker Compose**

创建 `docker-compose.yml`:
```yaml
version: '3.8'

services:
  web3-ai-agent:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    restart: unless-stopped
    volumes:
      - ./logs:/app/logs
```

2. **启动服务**
   ```bash
   docker-compose up -d
   ```

3. **查看日志**
   ```bash
   docker-compose logs -f
   ```

## 📁 项目结构

```
web3-ai-bsc-agent/
├── cmd/
│   └── main.go              # 应用入口
├── internal/
│   ├── api/
│   │   └── router.go        # HTTP 路由配置
│   ├── bsc/
│   │   ├── client.go        # BSC 客户端
│   │   ├── erc20.go         # ERC20 合约交互
│   │   ├── event.go         # 事件监听
│   │   └── tx.go            # 交易处理
│   └── ai/
│       └── agent.go         # AI 决策代理
├── .env                     # 环境变量配置
├── .gitignore
├── Dockerfile               # Docker 构建文件
├── go.mod
├── go.sum
└── README.md
```

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

### 开发规范

- 遵循 Go 代码规范
- 添加必要的注释（使用中文）
- 编写单元测试
- 更新文档

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🙏 致谢

- [go-ethereum](https://github.com/ethereum/go-ethereum) - Ethereum Go 实现
- [Gin](https://github.com/gin-gonic/gin) - Go Web 框架
- [OpenAI](https://openai.com/) - AI 能力支持

## 📞 联系方式

- GitHub: [@JackBean2134](https://github.com/JackBean2134)
- 项目地址: [https://github.com/JackBean2134/web3-ai-bsc-agent](https://github.com/JackBean2134/web3-ai-bsc-agent)

---

<div align="center">

如果这个项目对你有帮助，请给个 ⭐️ Star 支持一下！

</div>
