package gethclient

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"encoding/json"
)

type TransactionResult struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`             //"0xd7fb2ebd2037ff0959f6513343661772d279a7f5",
	Gas              string `json:"gas"`              //"0x88ce",
	GasPrice         string `json:"gasPrice"`         // "0x6fc23ac00",
	Hash             string `json:"hash"`             // "0x5d740a4a856206d51877568fe20405c9ca153a3165be5bab3ee84d1f1edd57c2",
	Input            string `json:"input"`            // "0xa9059cbb0000000000000000000000005ab6a67381a913fd071bd477148aed23b3c02c7a00000000000000000000000000000000000000000000000000000000000003e8",
	Nonce            string `json:"nonce"`            // "0x64",
	R                string `json:"r"`                // "0x1db6926d00d1021c60de31cd3f76dced61c2b0f7900c959943cad959cf689d53",
	S                string `json:"s"`                // "0x486704a7b6315aa052c4f141d7483f9279a2b780db9ca42295fad011aaadfdde",
	To               string `json:"to"`               // "0x38d2ab30612700e28eac03d7cc31fad2156ee274",
	TransactionIndex string `json:"transactionIndex"` // "0x0",
	V                string `json:"v"`                // "0x1b",
	Value            string `json:"value"`            // "0x0"
}

type TxPoolContent map[string]map[string]map[string]TransactionResult
type TxPoolInspect map[string]map[string]map[string]interface{}
type TxPoolInfo map[string]interface{}

func (tpc TxPoolContent) PrintPretty() {
	for status, pv := range tpc {
		fmt.Println("状态: ", status)
		for account, av := range pv {
			fmt.Println("account:", account)
			for nonceNum, transaction := range av {
				fmt.Printf("nonce:%v,transfaction:%#v\n", nonceNum, transaction)
			}
		}
	}
}

func (tpc TxPoolContent) Convert2json() string {
	bs, _ := json.Marshal(tpc)
	return string(bs)
}

func (tpc TxPoolInspect) PrintPretty() {
	for status, pv := range tpc {
		fmt.Println("状态: ", status)
		for account, av := range pv {
			fmt.Println("account:", account)
			for nonceNum, strSli := range av {
				fmt.Printf("nonce:%v,strSli:%v", nonceNum, strSli)
			}
		}
	}
}

func (tpc TxPoolInspect) Convert2json() string {
	bs, _ := json.Marshal(tpc)
	return string(bs)
}

func GetTxPoolContent() (txPoolContent TxPoolContent, err error) {
	txPoolContent = make(map[string]map[string]map[string]TransactionResult)
	err = Client.Call(&txPoolContent, "txpool_content")
	if err != nil {
		logs.Error("txpool_content failed, %v", err)
		return
	}
	return
}

func GetTxPoolInspect() (txPoolInspect TxPoolInspect, err error) {
	txPoolInspect = make(map[string]map[string]map[string]interface{})
	err = Client.Call(&txPoolInspect, "txpool_inspect")
	if err != nil {
		logs.Error("txpool_inspect failed, %v", err)
		return
	}
	return
}

func GetTxPoolInfo() (txPoolInfo TxPoolInfo, err error) {
	txPoolInfo = make(map[string]interface{})
	err = Client.Call(&txPoolInfo, "txpool_status")
	if err != nil {
		logs.Error("txpool_status failed, %v", err)
		return
	}
	return
}
