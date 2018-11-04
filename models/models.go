package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/Unknwon/com"
	"os"
	"path"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ethereum/go-ethereum/log"
	"errors"
	"time"
	"math/rand"
	"github.com/astaxie/beego/logs"
	"jifenblock/jifen"
)

const (
	_DB_NAME        = "data/jifen.db"
	_SQLITE3_DRIVER = "sqlite3"
	REMOVEJIFEN     = 0
	ADDJIFEN        = 1
)

type JifenTransaction struct {
	Id           int
	ContractAddr string
	TxHash       string
	Cost         int64
	GasCount     string
	GasPrice     string
	Nonce        string
	Value        int64
}

// 后期可以扩展Account字段(与前端用户做关联,username-->MemberOID--->积分)
type Account struct {
	Id         int64
	Name       string    `orm:"column(username);unique"`
	MemberOID  string    `orm:"column(memberid);unique"`
	CompanyOID string    `orm:"column(companyid);default(company)"`
	SiteOID    string    `orm:"column(siteid);default(site)"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
	Score      int64     `orm:"default(0)"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 需要在init中注册定义的model
	orm.RegisterModel(new(JifenTransaction), new(jifen.TxDetail), new(Account))
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddAccount(name, memberIdStr string) (err error) {
	o := orm.NewOrm()
	account := new(Account)
	account.Name = name
	account.MemberOID = memberIdStr
	if err := o.Read(account, "username"); err == nil {
		return errors.New("记录已经存在")
	}
	_, err = o.Insert(account)
	if err != nil {
		log.Debug("insert account failed name=%s, memberIdStr=%s, %v", name, memberIdStr, err)
		return err
	}
	return
}

func ModifyAccount(memberId string, balance int64) (err error) {
	o := orm.NewOrm()
	account := new(Account)
	account.MemberOID = memberId
	if err = o.Read(account, "memberid"); err != nil {
		return errors.New("ModifyAccount,不存在这个账户")
	}
	account.Score = balance
	_, err = o.Update(account)
	return
}

func GetAllAccounts() ([]*Account, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("account")
	topics := make([]*Account, 0)
	qs = qs.OrderBy("-created")
	_, err := qs.All(&topics)
	return topics, err
}

// txdetail
func AddTxDetail(detail *jifen.TxDetail) (err error) {
	o := orm.NewOrm()
	if err := o.Read(detail, "txid"); err == nil {
		return errors.New("tx_detail:记录已经存在")
	}
	_, err = o.Insert(detail)
	if err != nil {
		log.Debug("insert tx_detail failed txdetail=%v, %v", detail, err)
		return err
	}
	return
}

func GetAllTxDetails(memberId string) (txDetails []*jifen.TxDetail, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tx_detail")
	txDetails = make([]*jifen.TxDetail, 0)
	qs = qs.OrderBy("-createtime")
	if len(memberId) != 0 {
		qs = qs.Filter("memberid", memberId)
	}
	_, err = qs.All(&txDetails)
	return
}

/**
usercheck
 */
func CheckUserByName(username string) (err error) {
	o := orm.NewOrm()
	account := new(Account)
	account.Name = username

	if err := o.Read(account, "username"); err != nil {
		return errors.New(err.Error())
	}
	return
}
func CheckUserByMemberId(memberId string) (err error) {
	o := orm.NewOrm()
	account := new(Account)
	account.MemberOID = memberId

	if err := o.Read(account, "memberid"); err != nil {
		return errors.New(err.Error())
	}
	return
}

/**
生成memberID和txID
 */
func GetMemberIDStr() (memberIdStr string) {
	rand.Seed(time.Now().UnixNano())
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)
	memberSli := make([]byte, 0, 32)
	for i := 0; i < 32; i++ {
		memberSli = append(memberSli, bytes[rand.Intn(len(bytes))])
	}
	memberIdStr = string(memberSli)
	return
}

/**
生成交易信息
 */
func GenerateTxDetail(username string, txtype int, value int64) (txDetail jifen.TxDetail, err error) {
	o := orm.NewOrm()
	account := new(Account)
	account.Name = username

	if err = o.Read(account, "username"); err != nil {
		err = errors.New("不存在此账户")
		logs.Error("GenerateTxDetail failed, %v", err)
		return
	}

	tagq := "distribute"
	if txtype == REMOVEJIFEN {
		tagq = "deduct"
	}
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	txDetail = jifen.TxDetail{
		TxId:                                           GetMemberIDStr(),
		MemberOID:                                      account.MemberOID,
		TxType:                                         txtype,
		CompanyOID:                                     "xxxx",
		SiteOID:                                        "yyy",
		CreateTime:                                     timeStr,
		UpdateTime:                                     timeStr,
		Integral:                                       value,
		TpesDeductibleAmountConsumeAmountGoodsGoodsQty: tagq,
	}
	return
}
