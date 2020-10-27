package config

import "time"

type RpcConfig struct {
	DefaultFailover    string `yaml:"defaultFailover" json:"defaultFailover,omitempty" default:"failFast" property:"defaultFailover"`
	DefaultLoadBalance string `yaml:"defaultLoadBalance" json:"defaultLoadBalance,omitempty" default:"RoundRobin" property:"defaultLoadBalance"`
	Invoke_Timeout     string `yaml:"invoke_timeout" default:"3s" json:"invoke_timeout,omitempty" property:"invoke_timeout"`
	InvokeTimeout      time.Duration
	AppName            string `yaml:"appName" default:"unkonwn" json:"appName,omitempty" property:"appName"`

	*RegistryConfig
}
