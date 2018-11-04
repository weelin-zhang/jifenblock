package gethclient

import (
	"github.com/astaxie/beego/logs"
	"github.com/onrik/ethrpc"
	"math/big"
)

/**
获取该账户当前交易的nonce(pending, latest, earliest)
 */
func GetTransactionCount(accountAddr, status string) (count int, err error) {
	count, err = EthRpcClient.EthGetTransactionCount(accountAddr, status)
	if err != nil {
		logs.Error("GetTransactionCount failed, %v", err)
	}
	return
}

/**
发送以太币
	tx := ethrpc.T{
		From:     from,
		To:       to,
		//Gas:      30000,  //默认90000,多余退还
		GasPrice: big.NewInt(30000000000),
		Value:    big.NewInt(4947000000000000000),
	}
 */
func SendTransaction(tx ethrpc.T) (hash string, err error) {
	hash, err = EthRpcClient.EthSendTransaction(tx)
	if err != nil {
		logs.Error("SendTransaction failed, %v", err)
	}
	return
}

func GetWeb3ClientVersion() (version string, err error) {
	version, err = EthRpcClient.Web3ClientVersion()
	if err != nil {
		logs.Error("GetWeb3ClientVersion failed, %v", err)
	}
	return
}

/**
String - The current network protocol version
 */
func GetEthProtocolVersion() (version string, err error) {
	version, err = EthRpcClient.EthProtocolVersion()
	if err != nil {
		logs.Error("eth_protocolVersion failed, %v", err)
	}
	return
}

/**
Object|Boolean, An object with sync status data or FALSE, when not syncing:

startingBlock: QUANTITY - The block at which the import started (will only be reset, after the sync reached his head)
currentBlock: QUANTITY - The current block, same as eth_blockNumber
highestBlock: QUANTITY - The estimated highest block

eg:
&{false 0 0 0}||&{true z y x}

 */
func GetEthSyncing() (syncing *ethrpc.Syncing, err error) {
	syncing, err = EthRpcClient.EthSyncing()
	if err != nil {
		logs.Error("eth_protocolVersion failed, %v", err)
	}
	return
}

/**
矿工账号
 */
func GetEthCoinBase() (coinBase string, err error) {
	coinBase, err = EthRpcClient.EthCoinbase()
	if err != nil {
		logs.Error("eth_coinbase failed, %v", err)
	}
	return
}

/**
挖矿中..?
 */
func GetEthMining() (mining bool, err error) {
	mining, err = EthRpcClient.EthMining()
	if err != nil {
		logs.Error("eth_mining failed, %v", err)
	}
	return
}

/**
返回节点挖矿时每秒可算出的哈希数量
*/

func GetEthHashRate() (count int, err error) {
	count, err = EthRpcClient.EthHashrate()
	if err != nil {
		logs.Error("eth_hashrate failed, %v", err)
	}
	return
}

/**
eth_gasPrice
return: big.INT
 */

func GetEthGasPrice() (count big.Int, err error) {
	count, err = EthRpcClient.EthGasPrice()
	if err != nil {
		logs.Error("eth_gasPrice failed, %v", err)
	}
	return
}

/**

 */
func GetEthAccounts() (accounts []string, err error) {
	accounts, err = EthRpcClient.EthAccounts()
	if err != nil {
		logs.Error("eth_accounts failed, %v", err)
	}
	return
}

/**

 */
func GetEthBlockNumber() (num int, err error) {
	num, err = EthRpcClient.EthBlockNumber()
	if err != nil {
		logs.Error("eth_blockNumber failed, %v", err)
	}
	return
}

/**
 */

func GetEthBalance(addr string) (hexBalance string, err error) {
	hexBalance, err = EthRpcClient.EthGetBalance(addr, "pending")
	if err != nil {
		logs.Error("eth_getBalance failed, %v", err)
	}
	return
}

/**
Returns the number of transactions in a block from a block matching the given block hash.
根据block hash获取上面的交易数目
*/

func GetEthBlockTransactionCountByHash(hash string) (count int, err error) {
	count, err = EthRpcClient.EthGetBlockTransactionCountByHash(hash)
	if err != nil {
		logs.Error("eth_getBlockTransactionCountByHash failed, %v", err)
	}
	return
}

/**
Returns the number of transactions in a block from a block matching the given block hash.
根据block num获取上面的交易数目
 */
func GetEthBlockTransactionCountByNumber(num int) (count int, err error) {
	count, err = EthRpcClient.EthGetBlockTransactionCountByNumber(num)
	if err != nil {
		logs.Error("eth_getBlockTransactionCountByNumber failed, %v", err)
	}
	return
}

/**
Returns the number of uncles in a block from a block matching the given block number.
 */
func GetEthUncleCountByBlockNumber(num int) (count int, err error) {
	count, err = EthRpcClient.EthGetUncleCountByBlockNumber(num)
	if err != nil {
		logs.Error("eth_getUncleCountByBlockNumber failed, %v", err)
	}
	return
}

/**
Returns the number of uncles in a block from a block matching the given block hash.
 */
func GetEthUncleCountByBlockHash(hash string) (count int, err error) {
	count, err = EthRpcClient.EthGetUncleCountByBlockHash(hash)
	if err != nil {
		logs.Error("eth_getUncleCountByBlockHash failed, %v", err)
	}
	return
}

/**
Returns the number of uncles in a block from a block matching the given block number.
 */
func GetEthBlockByNumber(num int, withTx bool) (block *ethrpc.Block, err error) {
	block, err = EthRpcClient.EthGetBlockByNumber(num, withTx)
	if err != nil {
		logs.Error("eth_getBlockByNumber failed, %v", err)
	}
	return
}

/**
Returns the number of uncles in a block from a block matching the given block hash.
 */
func GetEthBlockByHash(hash string, withTx bool) (block *ethrpc.Block, err error) {
	block, err = EthRpcClient.EthGetBlockByHash(hash, withTx)
	if err != nil {
		logs.Error("eth_getBlockByHash failed, %v", err)
	}
	return
}

/**
Returns the information about a transaction requested by transaction hash.
 */
func GetEthTransactionByTxHash(hash string) (block *ethrpc.Transaction, err error) {
	block, err = EthRpcClient.EthGetTransactionByHash(hash)
	if err != nil {
		logs.Error("eth_getTransactionByHash failed, %v", err)
	}
	return
}

/**
Returns information about a transaction by *block* hash and transaction index position.
索引从1开始
 */
func GetEthTransactionByBlockHashAndIndex(hash string, num int) (block *ethrpc.Transaction, err error) {
	block, err = EthRpcClient.EthGetTransactionByBlockHashAndIndex(hash, num)
	if err != nil {
		logs.Error("eth_getTransactionByBlockHashAndIndex failed, %v", err)
	}
	return
}

/**
Returns information about a transaction by *block* hash and transaction index position.
索引从0开始
 */
func GetEthTransactionByBlockNumberAndIndex(blockNumber int, num int) (block *ethrpc.Transaction, err error) {
	block, err = EthRpcClient.EthGetTransactionByBlockNumberAndIndex(blockNumber, num)
	if err != nil {
		logs.Error("eth_getTransactionByBlockNumberAndIndex failed, %v", err)
	}
	return
}

/**
Returns information about a transaction by *block* hash and transaction index position.
索引从0开始
 */
func GetEthTransactionReceipt(hash string) (receipt *ethrpc.TransactionReceipt, err error) {
	receipt, err = EthRpcClient.EthGetTransactionReceipt(hash)
	if err != nil {
		logs.Error("eth_getTransactionReceipt failed, %v", err)
	}
	return
}
