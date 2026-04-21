package bsc

import (
	"context"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendBNB(client *ethclient.Client, to string, amountWei *big.Int) (*types.Transaction, error) {
	pk, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	from := crypto.PubkeyToAddress(pk.PublicKey)
	toAddr := common.HexToAddress(to)

	chainID := big.NewInt(97)
	nonce, _ := client.NonceAt(context.Background(), from, nil)

	tx := types.NewTransaction(nonce, toAddr, amountWei, 21000, big.NewInt(10e9), nil)
	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)

	return tx, client.SendTransaction(context.Background(), signedTx)
}
