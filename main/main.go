package main

import (
	"github.com/astaxie/beego/logs"
	"jifenblock/jifen"
	"jifenblock/models"
	"github.com/astaxie/beego"
	_ "jifenblock/routers"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init() {
	models.RegisterDB()
	orm.RunSyncdb("default", false, true)
}
func main() {
	err := loadConf()
	if err != nil {
		logs.Error("loadConf failed", err)
		panic(err)
	}

	err = initLogger("./logs/jifen.log", "debug")
	if err != nil {
		logs.Error("initLogger failed", err)
		panic(err)
	}

	err = jifen.InitJifen()
	if err != nil {
		logs.Error("initJifen failed", err)
		panic(err)
	}

	// 同步以太坊中最新积分数据
	go models.UpdateAccountTaskGoroutine()

	jifen.GetTxDetail("wRND6kYkpq3VaV81f1TSYY9UaPtrw6yh")
	fmt.Println(jifen.GetContactOwner())

	beego.Run()
}
