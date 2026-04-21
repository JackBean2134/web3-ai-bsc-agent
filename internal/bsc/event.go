package bsc

import (
	"context"
	"log"
	"math/big"
	"time"
	"web3-ai-agent/internal/ai"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ListenTransferEvent 监听ERC20转账事件并进行AI决策
func ListenTransferEvent(client *ethclient.Client, contract string) {
	if contract == "" {
		log.Fatal("CONTRACT_ADDR environment variable is not set")
	}

	addr := common.HexToAddress(contract)
	// ERC20 Transfer事件的topic签名
	topic := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
		Topics:    [][]common.Hash{{topic}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to filter logs: %v", err)
	}
	defer sub.Unsubscribe()

	log.Println("👂 Listening for ERC20 Transfer events...")

	for {
		select {
		case logData := <-logs:
			processTransferLog(logData)
		case err := <-sub.Err():
			log.Printf("Subscription error: %v, reconnecting in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			// 重新订阅
			sub, err = client.SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				log.Printf("Failed to reconnect: %v", err)
				continue
			}
			log.Println("Reconnected to event stream")
		}
	}
}

// processTransferLog 处理单个转账日志
func processTransferLog(logData types.Log) {
	if len(logData.Topics) < 3 {
		log.Println("Invalid log format: insufficient topics")
		return
	}

	from := common.BytesToAddress(logData.Topics[1].Bytes()).Hex()
	to := common.BytesToAddress(logData.Topics[2].Bytes()).Hex()
	value := new(big.Int).SetBytes(logData.Data)

	log.Printf("🔄 Transfer detected: %s -> %s | Amount: %s wei", from, to, value.String())

	// 调用AI进行风险评估
	decision, err := ai.AgentDecision(from, to, value.String())
	if err != nil {
		log.Printf("AI decision error: %v", err)
		return
	}

	log.Printf("🤖 AI Decision: %s", decision)

	// 根据AI决策执行相应操作
	handleAIDecision(decision, from, to, value)
}

// handleAIDecision 根据AI决策执行相应操作
func handleAIDecision(decision, from, to string, value *big.Int) {
	switch decision {
	case "safe":
		log.Printf("✅ Transaction deemed safe: %s -> %s", from, to)
		// 可以执行安全交易相关的操作
	case "risky":
		log.Printf("⚠️  Risky transaction detected: %s -> %s", from, to)
		// 可以执行风险控制相关的操作
	case "ignore":
		log.Printf("❌ Ignoring transaction: %s -> %s", from, to)
		// 忽略此交易
	default:
		log.Printf("❓ Unknown decision: %s for transaction %s -> %s", decision, from, to)
	}
}
