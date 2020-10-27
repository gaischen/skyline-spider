package config

import "sync"

var (
	rpcConfig   *RpcConfig
	configMutex sync.Mutex
)

func GetRpcConfig() *RpcConfig {
	if nil == rpcConfig {
		defer configMutex.Unlock()
		configMutex.Lock()
		return &RpcConfig{}
	}
	return rpcConfig
}
