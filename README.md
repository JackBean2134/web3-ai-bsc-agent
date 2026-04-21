# Web3 AI BSC Agent

<div align="center">

[![English](https://img.shields.io/badge/English-README-blue)](README.md)
[![中文](https://img.shields.io/badge/中文-README-red)](README.zh-CN.md)

![Go](https://img.shields.io/badge/Go-1.22+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Docker](https://img.shields.io/badge/Docker-Supported-orange.svg)
![Web3](https://img.shields.io/badge/Web3-Ethereum-purple.svg)

**A production-grade Web3 AI agent service with BSC chain event listening, ERC20 contract auto-parsing, AI risk decision-making, and automated on-chain operations**

</div>

## 📋 Table of Contents

- [Introduction](#-introduction)
- [Tech Stack](#-tech-stack)
- [Features](#-features)
- [Quick Start](#-quick-start)
- [API Documentation](#-api-documentation)
- [Deployment](#-deployment)
- [Project Structure](#-project-structure)
- [Contributing](#-contributing)
- [License](#-license)

## 🎯 Introduction

Web3 AI BSC Agent is an innovative blockchain intelligent agent system that combines the high performance of Go language, the low-cost advantage of BSC chain, and the decision-making capability of artificial intelligence. The system can monitor on-chain activities in real-time, automatically parse ERC20 token information, and conduct risk assessment and decision-making on transactions through AI models.

### Use Cases
- 🔍 **Risk Monitoring**: Real-time detection of suspicious transaction patterns
- 🤖 **Automated Trading**: Execute on-chain operations based on AI decisions
- 📊 **Data Analysis**: Collect and analyze on-chain data
- 🛡️ **Security Protection**: Identify potential fraudulent behavior

## 💻 Tech Stack

| Category | Technology |
|----------|------------|
| **Programming Language** | Go 1.22+ |
| **Web Framework** | Gin Framework |
| **Blockchain Interaction** | go-ethereum |
| **AI Integration** | OpenAI API / Compatible LLMs |
| **Containerization** | Docker |
| **Configuration Management** | Environment Variables |

## ✨ Features

- ✅ **BSC Chain Connection**: Stable and reliable BSC node connection
- ✅ **ERC20 Auto-Parsing**: Automatically retrieve token name, symbol, decimals, etc.
- ✅ **Real-time Event Listening**: Listen to Transfer events and respond in real-time
- ✅ **AI Risk Decision**: Transaction risk assessment based on large language models
- ✅ **RESTful API**: Complete HTTP API interface
- ✅ **Docker Deployment**: One-click containerized deployment
- ✅ **Error Handling**: Comprehensive error handling and logging
- ✅ **Graceful Shutdown**: Support smooth restart and shutdown

## 🚀 Quick Start

### Prerequisites

- Go 1.22+
- Docker (optional)
- BSC RPC node access
- OpenAI API Key or compatible LLM API

### Local Running

1. **Clone the repository**
   ```bash
   git clone https://github.com/JackBean2134/web3-ai-bsc-agent.git
   cd web3-ai-bsc-agent
   ```

2. **Configure environment variables**
   ```bash
   cp env.example .env
   # Edit .env file and fill in your configuration
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the service**
   ```bash
   go run cmd/main.go
   ```

5. **Access the API**
   ```
   http://localhost:8080/health
   ```

### Docker Running

```bash
# Build image
docker build -t web3-ai-bsc-agent .

# Run container
docker run -d \
  --name web3-ai-bsc-agent \
  -p 8080:8080 \
  --env-file .env \
  web3-ai-bsc-agent
```

## 📡 API Documentation

### Health Check

```bash
GET /health
```

**Response Example:**
```json
{
  "status": "ok",
  "message": "Web3 AI BSC Agent is running"
}
```

### Query BNB Balance

```bash
GET /balance/:address
```

**Parameters:**
- `address`: Ethereum address

**Response Example:**
```json
{
  "address": "0x1234567890abcdef1234567890abcdef12345678",
  "bnb_wei": "1000000000000000000"
}
```

### Query ERC20 Token Info

```bash
GET /erc20/info?contract=0x...
```

**Parameters:**
- `contract`: ERC20 contract address

**Response Example:**
```json
{
  "contract": "0x...",
  "name": "Binance USD",
  "symbol": "BUSD",
  "decimals": 18
}
```

### Query ERC20 Token Balance

```bash
GET /erc20/balance?contract=0x...&address=0x...
```

**Parameters:**
- `contract`: ERC20 contract address
- `address`: Wallet address

**Response Example:**
```json
{
  "contract": "0x...",
  "address": "0x...",
  "balance": "1000000000000000000"
}
```

### AI Trading Decision

```bash
GET /ai/decision?from=0x...&to=0x...&amount=100
```

**Parameters:**
- `from`: Sender address
- `to`: Receiver address
- `amount`: Transaction amount (optional)

**Response Example:**
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

## 🐳 Deployment Guide

### Environment Variable Configuration

Create `.env` file and configure the following variables:

```env
# BSC Node Configuration
BSC_RPC_URL=https://bsc-dataseed.binance.org/

# AI Configuration
OPENAI_API_KEY=your-api-key
OPENAI_MODEL=gpt-3.5-turbo

# Service Configuration
SERVER_PORT=8080
LOG_LEVEL=info

# Watched Contract Address (optional)
WATCH_CONTRACT=0x...
```

### Production Deployment

1. **Using Docker Compose**

Create `docker-compose.yml`:
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

2. **Start the service**
   ```bash
   docker-compose up -d
   ```

3. **View logs**
   ```bash
   docker-compose logs -f
   ```

## 📁 Project Structure

```
web3-ai-bsc-agent/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   └── router.go        # HTTP router configuration
│   ├── bsc/
│   │   ├── client.go        # BSC client
│   │   ├── erc20.go         # ERC20 contract interaction
│   │   ├── event.go         # Event listening
│   │   └── tx.go            # Transaction processing
│   └── ai/
│       └── agent.go         # AI decision agent
├── .env                     # Environment variable configuration
├── .gitignore
├── Dockerfile               # Docker build file
├── go.mod
├── go.sum
└── README.md
```

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go coding standards
- Add necessary comments (use English)
- Write unit tests
- Update documentation

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- [go-ethereum](https://github.com/ethereum/go-ethereum) - Ethereum Go implementation
- [Gin](https://github.com/gin-gonic/gin) - Go web framework
- [OpenAI](https://openai.com/) - AI capability support

## 📞 Contact

- GitHub: [@JackBean2134](https://github.com/JackBean2134)
- Project URL: [https://github.com/JackBean2134/web3-ai-bsc-agent](https://github.com/JackBean2134/web3-ai-bsc-agent)

---

<div align="center">

If this project helps you, please give it a ⭐️ Star!

</div>