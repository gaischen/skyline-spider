package rpc_client

import (
	"github.com/skyline/skyline-spider/rpc/framework"
	"github.com/skyline/skyline-spider/rpc/framework/entity"
	"github.com/skyline/skyline-spider/rpc/transport"
)

type ProviderInvoker struct {
}

func NewProviderInvoker() *ProviderInvoker {
	return &ProviderInvoker{}
}

func (p *ProviderInvoker) Invoke(context *framework.InvokerContext) (*entity.Result, error) {
	var err error
	var provider *transport.Provider



	return nil, err
}
