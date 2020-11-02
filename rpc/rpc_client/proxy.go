package rpc_client

import (
	"github.com/vanga-top/skyline-spider/rpc/framework"
	"github.com/vanga-top/skyline-spider/rpc/framework/config"
)

type proxy struct {
}

func newProxy(wrapper *rpcConsumerWrapper, rpcConfig *config.RpcConfig, invoker framework.Invoker) *proxy {

	return nil
}

func (p *proxy) doProxy() {

}
