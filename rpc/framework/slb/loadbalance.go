package slb

import (
	"github.com/skyline/skyline-spider/rpc/framework/config"
	"github.com/skyline/skyline-spider/rpc/framework/entity"
)

type LoadBalance interface {
	Select(providers []*config.URL, invocation *entity.Invocation) *config.URL
}

var (
	loadbalances = make(map[string]func() LoadBalance)
)

func GetLoadBalance(str string) LoadBalance {
	return loadbalances[str]()
}
