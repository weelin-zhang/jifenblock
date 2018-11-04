package gethclient

import "github.com/astaxie/beego/logs"


/**
newtwork id
 */
func GetNetVersion() (version string, err error) {
	version, err = EthRpcClient.NetVersion()
	if err != nil {
		logs.Error("net_version failed, %v", err)
	}
	return
}

/**
net_peerCount
 */
func GetPeerCount() (count int, err error) {
	count, err = EthRpcClient.NetPeerCount()
	if err != nil {
		logs.Error("net_peerCount failed, %v", err)
	}
	return
}

/**
net listening
 */
func GetNetListening() (listen bool , err error) {
	listen, err = EthRpcClient.NetListening()
	if err != nil {
		logs.Error("net_listening failed, %v", err)
	}
	return
}
