package gethclient

import (
	"github.com/astaxie/beego/logs"
	"fmt"
)

/**
Imports the given unencrypted private key (hex string) into the key store, encrypting it with the passphrase.
通过私钥(hex string)和passphrase 把账号导入key store
*/
func GetAccountAddr(privateKey, passphrase string) (accountAddr string, err error) {
	err = Client.Call(&accountAddr, "personal_importRawKey", privateKey, passphrase)
	if err != nil {
		logs.Error("personal_importRawKey failed, %v", err)
		return
	}
	return
}

func GetAllAccounts() (accounts []string, err error) {
	err = Client.Call(&accounts, "personal_listAccounts")
	if err != nil {
		logs.Error("getAllAccounts failed, %v", err)
		return
	}
	return
}

func NewAccount(passphrase string) (account string, err error) {
	err = Client.Call(&account, "personal_newAccount", passphrase)
	if err != nil {
		logs.Error("newAccount failed, %v", err)
		return
	}
	return
}

func LockAccount(accountAddr string) (lock bool, err error) {
	err = Client.Call(&lock, "personal_lockAccount", accountAddr)
	if err != nil {
		logs.Error("personal_lockAccount failed, %v", err)
		return
	}
	return
}

/**
druation: 解锁有效时间, default 300s
 */
func UnLockAccount(accountAddr, passphrase string, duration int) (unLock bool, err error) {
	err = Client.Call(&unLock, "personal_unlockAccount", accountAddr, passphrase, duration)
	fmt.Println(unLock)
	if err != nil {
		logs.Error("personal_unlockAccount failed, %v", err)
		return
	}
	return
}
