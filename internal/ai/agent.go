package ai

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

// AgentDecision 使用AI分析交易的安全性
func AgentDecision(from, to, amount string) (string, error) {
	apiKey := os.Getenv("LLM_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("LLM_API_KEY environment variable is not set")
	}

	client := openai.NewClient(apiKey)

	prompt := fmt.Sprintf(`你是Web3安全AI助手，请分析以下交易的风险等级。
仅返回以下三种结果之一: safe / risky / ignore

交易详情:
- 发送方: %s
- 接收方: %s
- 金额: %s wei

请基于地址模式和金额大小判断风险等级。`, from, to, amount)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned from AI")
	}

	decision := resp.Choices[0].Message.Content
	log.Printf("AI Decision Result: %s", decision)

	return decision, nil
}
