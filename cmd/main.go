package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"web3-ai-agent/internal/api"
	"web3-ai-agent/internal/bsc"

	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// 验证必要的环境变量
	requiredEnvVars := []string{"RPC_URL", "PRIVATE_KEY", "CONTRACT_ADDR", "LLM_API_KEY"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Required environment variable %s is not set", envVar)
		}
	}

	// 初始化 Web3 客户端
	client, err := bsc.NewEthClient(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to BSC node: %v", err)
	}
	defer client.Close()

	log.Println("Successfully connected to BSC node")

	// 启动事件监听（后台协程）
	go func() {
		bsc.ListenTransferEvent(client, os.Getenv("CONTRACT_ADDR"))
	}()

	// 设置信号处理，优雅关闭
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 启动 HTTP API
	r := api.SetupRouter(client)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverAddr := ":" + port
	log.Printf("Starting HTTP server on %s", serverAddr)

	// 在goroutine中启动服务器，以便我们可以等待中断信号
	go func() {
		if err := r.Run(serverAddr); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// 等待中断信号
	<-ctx.Done()

	log.Println("Shutting down gracefully...")
	stop()
	log.Println("Server stopped")
}
