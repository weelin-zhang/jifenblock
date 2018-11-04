package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"jifenblock/models"
	"github.com/astaxie/beego/logs"
	"jifenblock/jifen"
	"strconv"
	"fmt"
	"time"
)

type JifenController struct {
	beego.Controller
}

func (c *JifenController) Home() {
	fmt.Println(c.Ctx.Request.RequestURI)
	c.Ctx.WriteString(c.Ctx.Request.RemoteAddr)

}

func (c *JifenController) Dashboard() {
	c.Data["ContentTitle"] = "Dashboard"
	c.Data["BreadCrumb"] = template.HTML(`<li class="active">Dashboard</li>`)
	c.Data["IsDashboard"] = true
	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "base/header.html"
	c.LayoutSections["Sidebar"] = "base/sidebar.html"
	c.TplName = "dashboard.html"
	accounts, _ := models.GetAllAccounts()
	c.Data["AccountNum"] = len(accounts)
}

func (c *JifenController) Accounts() {
	c.Data["Accounts"] = jifen.JifenConfig.AccountList
	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "base/header.html"
	c.LayoutSections["Sidebar"] = "base/sidebar.html"
	c.LayoutSections["Scripts"] = "account/accounts_js.html"
	c.TplName = "account/accounts.html"

	c.Data["ContentTitle"] = "用户"
	c.Data["BreadCrumb"] = template.HTML(`<li class="active">Accounts</li>`)
	c.Data["IsAccounts"] = true

	accounts, err := models.GetAllAccounts()
	if err != nil {
		logs.Error("GetAllAccounts failed, %v", err)
		return
	}
	// 账户score同步, memberId传入chan
	//go func() {
	//	for _, account := range accounts {
	//		jifen.JifenConfig.NeedUpdateAccount <- account.MemberOID
	//	}
	//}()

	c.Data["Accounts"] = accounts
}

/**
积分管理
 */

func (c *JifenController) ScoreManage() {
	username := c.Ctx.Input.Param(":username")

	if c.Ctx.Request.Method == "GET" {
		c.Layout = "base/layout.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Header"] = "base/header.html"
		c.LayoutSections["Sidebar"] = "base/sidebar.html"
		c.LayoutSections["Scripts"] = "account/accountmanage_js.html"
		c.TplName = "account/accountmanage.html"
		c.Data["BreadCrumb"] = template.HTML(`<li class="active">scoremanage</li>`)
		c.Data["CurAccountName"] = username
		return
	}
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "success"
	defer func() {
		c.Data["json"] = result
		c.ServeJSON()
	}()

	optype := c.GetString("optype")
	countStr := c.GetString("count")
	logs.Debug("optype:%s,username:%s, count: %s", optype, username, countStr)
	if len(optype) == 0 || len(countStr) == 0 {
		result["msg"] = "confirm input optype and count"
		result["code"] = 1001
		return
	}

	count, err := strconv.ParseInt(countStr, 10, 64)
	if err != nil {
		result["msg"] = "invalid params," + err.Error()
		result["code"] = 1002
		return
	}

	// 生成交易记录
	txDetail, err := models.GenerateTxDetail(username, models.ADDJIFEN, count)
	if optype == "deduct" {
		txDetail, err = models.GenerateTxDetail(username, models.REMOVEJIFEN, count)
	}
	if err != nil { // 积分合约中没有对应的账户
		result["msg"] = "invalid params," + err.Error()
		result["code"] = 1002
		return
	}

	switch optype {
	case "distribute":
		// 先把交易提交到以太坊/
		tx, err := jifen.DoDistributeJiFen(txDetail)
		fmt.Printf("交易信息:::::%T, %v, %v\n", tx, tx, err)
		// 检查交易是否成功
		if err != nil {
			logs.Error("交易执行失败 txDetail: %v,err: %v", txDetail, err)
			result["msg"] = "交易执行失败,err:" + err.Error()
			result["code"] = 1003
			return
		}
	case "deduct":
		// 先把交易提交到以太坊/
		tx, err := jifen.DoDeductJiFen(txDetail)
		fmt.Printf("交易信息:::::%T, %v, %v\n", tx, tx, err)

		// 保存or 返回
		if err != nil {
			logs.Error("交易执行失败 txDetail: %v,err: %v", txDetail, err)
			result["msg"] = "交易执行失败,err:" + err.Error()
			result["code"] = 1003
			return
		}
	}
	// 把待更新账户存入chan
	go func() {
		// 延时，等待交易进入penging状态
		time.Sleep(50 * time.Millisecond)
		jifen.JifenConfig.NeedUpdateAccount <- txDetail.MemberOID
	}()

	// 正常保存交易记录
	err = models.AddTxDetail(&txDetail)
	if err != nil {
		logs.Error("txDetail save failed, err: %v", err)
		result["msg"] = "txDetail save failed, err:" + err.Error()
		result["code"] = 1004
		return
	}

	//返回
	result["data"] = txDetail
	return
}

/**
刷新积分
*/

func (c *JifenController) RefreshScore() {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "success"
	defer func() {
		c.Data["json"] = result
		c.ServeJSON()
	}()

	memberId := c.Ctx.Input.Param(":memberId")
	if len(memberId) == 0 {
		result["msg"] = "no memberId"
		result["code"] = 1002
		return
	}

	// 更新account score
	balance, err := models.UpdateAccountScore(memberId)

	if err != nil {
		logs.Error("get account[%s], failed,err:%v", memberId, err)
		result["msg"] = "refresh  account balance failed err:" + err.Error()
		result["code"] = 1005
		return
	}
	result["data"] = struct {
		Balance int64
	}{Balance: balance}
}

/**
创建用户
 */
func (c *JifenController) CreateAccount() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["ErrMsg"] = false
		if errMsg := c.GetString("errMsg"); len(errMsg) > 0 {
			c.Data["ErrMsg"] = errMsg
		}
		c.Data["ContentTitle"] = "Createaccount"
		c.Data["BreadCrumb"] = template.HTML(`<li class="active">Createaccount</li>`)
		c.Data["IsCreateAccount"] = true
		c.Layout = "base/layout.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Header"] = "base/header.html"
		c.LayoutSections["Sidebar"] = "base/sidebar.html"
		c.LayoutSections["Scripts"] = "account/accounts_js.html"
		c.TplName = "account/createaccount.html"
		return
	} else {
		name := c.GetString("username")
		if len(name) == 0 {
			c.Redirect("/createaccount?errMsg=输入正确的用户名", 302)
		}
		// 获得对应的memberoid
		memberIdStr := models.GetMemberIDStr()
		err := models.AddAccount(name, memberIdStr)
		if err != nil {

			logs.Error("createAccount failed, name=%s, err:%v", name, err)
			c.Redirect("/createaccount?errMsg="+err.Error(), 302)
		}
		c.Redirect("/accounts", 302)
	}
}

/**
交易信息
 */
func (c *JifenController) Transactions() {
	c.Data["ContentTitle"] = "交易信息"
	c.Data["BreadCrumb"] = template.HTML(`<li class="active">Transactions</li>`)
	c.Data["IsTrans"] = true
	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "base/header.html"
	c.LayoutSections["Sidebar"] = "base/sidebar.html"
	c.TplName = "transactions.html"

	memberId := c.Ctx.Input.Param(":memberId")
	var txDetails []*jifen.TxDetail
	if len(memberId) == 0 {
		txs, err := models.GetAllTxDetails("")
		txDetails = txs
		if err != nil {
			logs.Error("get allTxDetails failed, %v", err)
			txDetails = []*jifen.TxDetail{}
		}
	} else {
		err := models.CheckUserByMemberId(memberId)
		if err != nil {
			logs.Error("CheckUserByMemberId failed, %v", err)
			c.Redirect("/dashboard", 302)
		}

		txs, err := models.GetAllTxDetails(memberId)
		txDetails = txs
		if err != nil {
			logs.Error("get allTxDetails failed,memberId:%s, err:%v", memberId, err)
			txDetails = []*jifen.TxDetail{}
		}
	}
	c.Data["TxDetails"] = txDetails
}
