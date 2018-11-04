package jifen

import (
	"strings"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"fmt"
	"github.com/astaxie/beego/logs"
)

/**
根据账户获取私钥文件
 */
func GetKeyFile(addr, keydir string) (string, error) {

	rds, err := ioutil.ReadDir(keydir)
	for _, rd := range rds {
		if !rd.IsDir() && strings.HasSuffix(rd.Name(), strings.ToLower(string(addr[2:]))) {
			return keydir + rd.Name(), err
		}
	}
	return "", err
}

/**
获取连接
 */
func GetCli(url string) (cli *ethclient.Client, err error) {
	// 建立连接
	cli, err = ethclient.Dial(url)
	return cli, err
}

/**
打印交易信息
 */
func PrintTx(tx *types.Transaction, params ...interface{}) {
	if len(params) > 0 {
		fmt.Printf("transfer pending,交易信息如下:\n 交易hash: %s, 花费(wei):%d, gas数量:%d, gas单价:%d,nonce:%d, value(amount): %d \n", tx.Hash().Hex(), tx.Cost().Uint64(), tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce(), tx.Value())
		logs.Info("transfer pending,交易信息如下:\n 交易hash: %s, 花费(wei):%d, gas数量:%d, gas单价:%d,nonce:%d, value(amount): %d \n", tx.Hash().Hex(), tx.Cost().Uint64(), tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce(), tx.Value())
	} else {
		fmt.Printf("transfer pending,交易信息如下:\n执行合约: %s, 交易hash: %s, 花费(wei):%d, gas数量:%d, gas单价:%d,nonce:%d, value(amount): %d \n", tx.To().Hex(), tx.Hash().Hex(), tx.Cost().Uint64(), tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce(), tx.Value())
		logs.Info("transfer pending,交易信息如下:\n执行合约: %s, 交易hash: %s, 花费(wei):%d, gas数量:%d, gas单价:%d,nonce:%d, value(amount): %d \n", tx.To().Hex(), tx.Hash().Hex(), tx.Cost().Uint64(), tx.Gas(), tx.GasPrice().Uint64(), tx.Nonce(), tx.Value())
	}
}

/**
生成32字符串-->[32]byte
 */
func ChangeStringTo32byte(str string) (arr [32]byte) {
	sli := arr[:]
	bytes := []byte(str)
	for i := 0; i < len(arr); i++ {
		sli[i] = bytes[i]
	}
	return
}
