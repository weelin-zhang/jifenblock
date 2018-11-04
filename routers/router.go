package routers

import (
	"github.com/astaxie/beego"
	"jifenblock/controllers"
)

func init() {
	beego.Router("/", &controllers.JifenController{}, "*:Home")
	beego.Router("/dashboard", &controllers.JifenController{}, "*:Dashboard")
	beego.Router("/accounts", &controllers.JifenController{}, "*:Accounts")
	//
	beego.Router("/createaccount", &controllers.JifenController{}, "*:CreateAccount")
	beego.Router("/scoremanage/:username", &controllers.JifenController{}, "*:ScoreManage")
	beego.Router("/refreshscore/:memberId", &controllers.JifenController{}, "*:RefreshScore")

	// 交易详情
	beego.Router("/transactions/?:memberId", &controllers.JifenController{}, "*:Transactions")
}
