package bsc

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// NewEthClient 创建并返回一个以太坊客户端连接
func NewEthClient(url string) (*ethclient.Client, error) {
	if url == "" {
		log.Fatal("RPC_URL environment variable is not set")
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	// 测试连接是否有效
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = client.ChainID(ctx)
	if err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}
