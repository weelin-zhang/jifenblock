package gethclient

import (
	"github.com/onrik/ethrpc"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/astaxie/beego/logs"
)

var Client *rpc.Client
var EthRpcClient *ethrpc.EthRPC

func init() {
	//rpc client
	c, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		logs.Error("rpc.Client init failed, %v", err)
		panic(err)
	}
	Client = c

	//ethrpc
	ethrpc := ethrpc.New("http://127.0.0.1:8545")
	EthRpcClient = ethrpc
}
