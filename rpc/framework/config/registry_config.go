package config

type RegistryConfig struct {
	ReferCheck       bool `yaml:"referCheck" default:"true" json:"referCheck,omitempty" property:"ReferCheck"`                   //启动时是否检查服务是否订阅
	SubscribeTimeout int64  `yaml:"subscribeTimeout" default:"6000" json:"subscribeTimeout,omitempty" property:"subscribeTimeout"` //服务订阅超时
}
