package jifen

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"github.com/astaxie/beego/logs"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"fmt"
)

func init() {
	JifenConfig = &JifenConf{
		NeedUpdateAccount: make(chan string, 100),
	}
}

/**
初始化jifen
 */

func InitJifen() (err error) {
	err = InitJifenSession()
	return
}

/**
初始化JifenSession
 */
func InitJifenSession() (err error) {
	cli, err := GetCli(JifenConfig.HostUrl)
	if err != nil {
		fmt.Println("InitJifenSession Failed, ", err)
		panic(err)
	}
	token, err := NewJifen(common.HexToAddress(JifenConfig.ContractAddr), cli)

	keyFilePath, err := GetKeyFile(JifenConfig.OwnerAccount, JifenConfig.KeyDir)
	if err != nil {
		fmt.Println("InitJifenSession Failed, ", err)
		return
	}
	fp, err := os.Open(keyFilePath)
	if err != nil {
		fmt.Println("InitJifenSession Failed, ", err)
		return
	}
	auth, err := bind.NewTransactor(fp, JifenConfig.OwerSecret)
	if err != nil {
		logs.Error("InitJifenSession Failed, cannot create auth %v", err)
		return
	}
	JifenConfig.JifenSession = &JifenSession{
		Contract: token,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:   auth.From,
			Signer: auth.Signer,
		},
	}
	return
}

/**
部署合约
 */
func DeployContract(contractOwnerAddr, passwd string, params ...interface{}) (tx *types.Transaction, contractAddr string, err error) {
	keyfilenamepath, err := GetKeyFile(contractOwnerAddr, JifenConfig.KeyDir)
	fp, err := os.Open(keyfilenamepath)
	if err != nil {
		return
	}
	defer fp.Close()

	auth, err := bind.NewTransactor(fp, passwd)
	if err != nil {
		return
	}

	backend, err := GetCli(JifenConfig.HostUrl)
	if err != nil {
		return
	}

	defer backend.Close()
	addr, tx, _, err := DeployJifen(auth, backend)
	contractAddr = addr.Hex()
	fmt.Println("合约部署合约地址: ", addr.Hex())
	PrintTx(tx, "deploy contract")
	return
}

/**
发放积分
 */

func DoDistributeJiFen(detail TxDetail) (tx *types.Transaction, err error) {
	value := big.NewInt(detail.Integral)
	txid := ChangeStringTo32byte(detail.TxId)
	memberOID := ChangeStringTo32byte(detail.MemberOID)
	companyOID := detail.CompanyOID
	siteOID := detail.SiteOID
	createTime := detail.CreateTime
	updateTime := detail.UpdateTime
	integral := string(detail.Integral)
	types_deductibleAmount_consumeAmount_goods_goodsQty := "distribute"
	tx, err = JifenConfig.JifenSession.DistributeJiFenFlow(value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
	if err != nil {
		logs.Error("积分发放failed, %v", err)
		return
	}
	PrintTx(tx)
	return nil, nil
}

/**
扣除积分
 */
func DoDeductJiFen(detail TxDetail) (tx *types.Transaction, err error) {
	value := big.NewInt(detail.Integral)
	txid := ChangeStringTo32byte(detail.TxId)
	memberOID := ChangeStringTo32byte(detail.MemberOID)
	companyOID := detail.CompanyOID
	siteOID := detail.SiteOID
	createTime := detail.CreateTime
	updateTime := detail.UpdateTime
	integral := string(detail.Integral)
	types_deductibleAmount_consumeAmount_goods_goodsQty := "deduct"
	tx, err = JifenConfig.JifenSession.DeductJiFenFlow(value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
	if err != nil {
		logs.Error("积分扣除failed, %v", err)
		return
	}
	PrintTx(tx)
	return nil, nil
}

/**
从以太坊获取交易详情
 */
func GetTxDetail(txId string) {
	txDetail, err := JifenConfig.JifenSession.TxDetails(ChangeStringTo32byte(txId))
	if err != nil {
		logs.Error(err)
		return
	}
	fmt.Println(string(txDetail.MemberOID[:]))
	fmt.Println(string(txDetail.SiteOID))
	fmt.Println(string(txDetail.CompanyOID))
	fmt.Printf("%q\n", txDetail.TypesDeductibleAmountConsumeAmountGoodsGoodsQty)
}

/**
获取账户积分数
 */
func GetAccountBalance(memberId string) (balance int64, err error) {
	b, err := JifenConfig.JifenSession.BalanceOf(ChangeStringTo32byte(memberId))
	if err != nil {
		logs.Error(err)
		return
	}
	balance = b.Int64()
	return
}

/**
获取用户交易次数
 */
func GetUserTxNumber(memberId string) (count int64, err error) {
	big, err := JifenConfig.JifenSession.GetUserTxsLength(ChangeStringTo32byte(memberId))
	if err != nil {
		logs.Error(err)
		return
	}
	count = big.Int64()
	return
}

/**
获取Owner
 */

func GetContactOwner() (owner string, err error) {
	tmpOwnerAddr, err := JifenConfig.JifenSession.Owner()
	if err != nil {
		logs.Error(err)
		return
	}
	owner = tmpOwnerAddr.String()
	return
}
