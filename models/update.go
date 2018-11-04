package models

import (
	"github.com/astaxie/beego/logs"
	"jifenblock/jifen"
	"fmt"
)

func UpdateAccountScore(memberId string) (balance int64, err error) {
	balance, err = jifen.GetAccountBalance(memberId)
	if err != nil {
		logs.Error("get account[%s], failed,err:%v", memberId, err)
		return
	}
	// 更新account
	err = ModifyAccount(memberId, balance)
	if err != nil {
		logs.Error("ModifyAccount get account[%s], failed,err:%v", memberId, err)
		return
	}
	return
}

func UpdateAccountTaskGoroutine() {
	for {
		memberId, ok := <-jifen.JifenConfig.NeedUpdateAccount
		if ok {
			logs.Info("NeedUpdateAccount remind count:", len(jifen.JifenConfig.NeedUpdateAccount))
			_, err := UpdateAccountScore(memberId)
			if err != nil {
				logs.Error("UpdateAccountScore by chan failed, memberId=%v", memberId)
				continue
			}
			fmt.Println("update...syning....")
		}
	}
}
