package bsc

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ERC20ABI 定义ERC20标准合约的ABI接口
const ERC20ABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

// erc20ABI 解析后的ABI对象
var erc20ABI, _ = abi.JSON(strings.NewReader(ERC20ABI))

// GetERC20Info 获取ERC20代币的基本信息
func GetERC20Info(client *ethclient.Client, contract string) (name, symbol string, decimals uint8, err error) {
	addr := common.HexToAddress(contract)
	ctx := context.Background()

	// 获取代币名称
	resName, err := client.CallContract(ctx, ethereum.CallMsg{To: &addr, Data: erc20ABI.Methods["name"].ID}, nil)
	if err != nil {
		return "", "", 0, err
	}
	err = erc20ABI.Methods["name"].Outputs.Unpack(&name, resName)
	if err != nil {
		return "", "", 0, err
	}

	// 获取代币符号
	resSymbol, err := client.CallContract(ctx, ethereum.CallMsg{To: &addr, Data: erc20ABI.Methods["symbol"].ID}, nil)
	if err != nil {
		return "", "", 0, err
	}
	err = erc20ABI.Methods["symbol"].Outputs.Unpack(&symbol, resSymbol)
	if err != nil {
		return "", "", 0, err
	}

	// 获取小数位数
	resDec, err := client.CallContract(ctx, ethereum.CallMsg{To: &addr, Data: erc20ABI.Methods["decimals"].ID}, nil)
	if err != nil {
		return "", "", 0, err
	}
	err = erc20ABI.Methods["decimals"].Outputs.Unpack(&decimals, resDec)
	if err != nil {
		return "", "", 0, err
	}

	return name, symbol, decimals, nil
}

// GetERC20Balance 获取指定地址的ERC20代币余额
func GetERC20Balance(client *ethclient.Client, contract, addr string) (*big.Int, error) {
	contractAddr := common.HexToAddress(contract)
	address := common.HexToAddress(addr)

	data, err := erc20ABI.Pack("balanceOf", address)
	if err != nil {
		return nil, err
	}

	res, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	err = erc20ABI.Methods["balanceOf"].Outputs.Unpack(&balance, res)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
