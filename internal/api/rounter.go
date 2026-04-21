package api

import (
	"log"
	"net/http"
	"web3-ai-bsc-agent/internal/ai"
	"web3-ai-bsc-agent/internal/bsc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置HTTP路由
func SetupRouter(client *ethclient.Client) *gin.Engine {
	r := gin.Default()

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Web3 AI BSC Agent is running"})
	})

	// 查询BNB余额
	r.GET("/balance/:address", func(c *gin.Context) {
		addr := c.Param("address")

		// 验证地址格式
		if !common.IsHexAddress(addr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Ethereum address"})
			return
		}

		bal, err := client.BalanceAt(c.Request.Context(), common.HexToAddress(addr), nil)
		if err != nil {
			log.Printf("Error fetching balance for %s: %v", addr, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch balance"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"address": addr, "bnb_wei": bal.String()})
	})

	// 查询ERC20代币信息
	r.GET("/erc20/info", func(c *gin.Context) {
		contract := c.Query("contract")

		if contract == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "contract address is required"})
			return
		}

		if !common.IsHexAddress(contract) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid contract address"})
			return
		}

		name, symbol, decimals, err := bsc.GetERC20Info(client, contract)
		if err != nil {
			log.Printf("Error fetching ERC20 info for %s: %v", contract, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch ERC20 info"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"contract": contract,
			"name":     name,
			"symbol":   symbol,
			"decimals": decimals,
		})
	})

	// 查询ERC20代币余额
	r.GET("/erc20/balance", func(c *gin.Context) {
		contract := c.Query("contract")
		addr := c.Query("address")

		if contract == "" || addr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "both contract and address parameters are required"})
			return
		}

		if !common.IsHexAddress(contract) || !common.IsHexAddress(addr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid contract or address parameter"})
			return
		}

		bal, err := bsc.GetERC20Balance(client, contract, addr)
		if err != nil {
			log.Printf("Error fetching ERC20 balance for %s at %s: %v", contract, addr, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch ERC20 balance"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"contract": contract,
			"address":  addr,
			"balance":  bal.String(),
		})
	})

	// AI决策端点
	r.GET("/ai/decision", func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")
		amount := c.Query("amount")

		if from == "" || to == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "from and to parameters are required"})
			return
		}

		if !common.IsHexAddress(from) || !common.IsHexAddress(to) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from or to address"})
			return
		}

		dec, err := ai.AgentDecision(from, to, amount)
		if err != nil {
			log.Printf("Error getting AI decision: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get AI decision"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"from":     from,
			"to":       to,
			"amount":   amount,
			"decision": dec,
		})
	})

	return r
}
