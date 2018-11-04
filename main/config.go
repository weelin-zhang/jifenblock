package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"jifenblock/jifen"
	"strings"
)

func loadConf() (err error) {
	// log
	jifen.JifenConfig.LogPath = beego.AppConfig.String("log_path")
	jifen.JifenConfig.LogLevel = beego.AppConfig.String("log_level")
	if len(jifen.JifenConfig.LogLevel) == 0 || len(jifen.JifenConfig.LogPath) == 0 {
		logs.Error("loadConf failed, log_path=%s, log_level=%s", jifen.JifenConfig.LogPath, jifen.JifenConfig.LogLevel)
		err = fmt.Errorf("loadConf failed log_path=%s, log_level=%s", jifen.JifenConfig.LogPath, jifen.JifenConfig.LogLevel)
		return
	}

	// accounts
	ownerSecret := beego.AppConfig.String("owner_secret")
	if len(ownerSecret) == 0 {
		logs.Error("loadConf failed, owner_secret=%s", ownerSecret)
		err = fmt.Errorf("loadConf failed ownr_secret", ownerSecret)
		return
	}
	jifen.JifenConfig.OwerSecret = ownerSecret

	ownerAccount := beego.AppConfig.String("owner_account")
	if len(ownerAccount) == 0 {
		logs.Error("loadConf failed, owner_account=%s", ownerAccount)
		err = fmt.Errorf("loadConf failed ownr_account", ownerAccount)
		return
	}
	jifen.JifenConfig.OwnerAccount = ownerAccount

	keyDir := beego.AppConfig.String("key_dir")
	if len(keyDir) == 0 {
		logs.Error("loadConf failed, key_dir=%s", keyDir)
		err = fmt.Errorf("loadConf failed key_dir", keyDir)
		return
	}
	jifen.JifenConfig.KeyDir = keyDir

	contractAddr := beego.AppConfig.String("contract_addr")
	if len(contractAddr) == 0 {
		logs.Error("loadConf failed, contract_addr=%s", contractAddr)
		err = fmt.Errorf("loadConf failed contract_addr", contractAddr)
		return
	}
	jifen.JifenConfig.ContractAddr = contractAddr

	hostAddr := beego.AppConfig.String("host_addr")
	if len(hostAddr) == 0 {
		logs.Error("loadConf failed, host_addr=%s", hostAddr)
		err = fmt.Errorf("loadConf failed host_addr", hostAddr)
		return
	}
	jifen.JifenConfig.HostUrl = hostAddr

	accountsAddr := beego.AppConfig.String("accounts_addr")
	if len(accountsAddr) == 0 {
		logs.Error("loadConf failed, accounts_addr=%s", accountsAddr)
		err = fmt.Errorf("loadConf failed accounts_addr", accountsAddr)
		return
	}
	jifen.JifenConfig.AccountList = strings.Split(accountsAddr, ",")
	return
}
