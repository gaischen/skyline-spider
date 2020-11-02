package rpc_client

import (
	"github.com/vanga-top/skyline-spider/rpc/framework"
	"github.com/vanga-top/skyline-spider/rpc/framework/entity"
	"github.com/vanga-top/skyline-spider/rpc/transport"
)

type ProviderInvoker struct {
}

func NewProviderInvoker() *ProviderInvoker {
	return &ProviderInvoker{}
}

func (p *ProviderInvoker) Invoke(context *framework.InvokerContext) (*entity.Result, error) {
	var err error
	var provider *transport.Provider

	provider = context.Provider
	if provider == nil {
		return &entity.Result{
			Status:      0,
			Value:       nil,
			Desc:        "",
			Exception:   nil,
			Attachments: nil,
		}, err
	}



	return nil, err
}
