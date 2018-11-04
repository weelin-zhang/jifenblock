package gethclient

import (
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/ethereum/go-ethereum/common/hexutil"
)


/**
设置矿工最小的可以接受的gasPrice
 */

func SetMinGasPrice(num int64) (err error) {
	err = Client.Call(nil, "miner_setGasPrice", hexutil.Big(*big.NewInt(num)))
	if err != nil {
		logs.Error("SetMinGasPrice failed, %v", err)
		return
	}
	return
}

/**
设置etherbase(矿工的收益账号)
*/
func SetEtherbase(coinBaseAddr string) (err error) {
	err = Client.Call(nil, "miner_setEtherbase", common.HexToAddress(coinBaseAddr))
	if err != nil {
		logs.Error("SetEtherbase failed, %v", err)
		return
	}
	return
}

/**
开始挖矿
*/
func StartMining() (err error) {
	err = Client.Call(nil, "miner_start", 1)
	if err != nil {
		logs.Error("startMining failed, %v", err)
		return
	}
	return
}

/**
结束挖矿
 */
func StopMining() (err error) {
	err = Client.Call(nil, "miner_stop")
	if err != nil {
		logs.Error("startMining failed, %v", err)
		return
	}
	return
}


