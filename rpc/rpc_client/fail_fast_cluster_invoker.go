package rpc_client

import (
	"github.com/vanga-top/skyline-spider/rpc/framework"
	"github.com/vanga-top/skyline-spider/rpc/framework/config"
	"github.com/vanga-top/skyline-spider/rpc/framework/entity"
	"github.com/vanga-top/skyline-spider/rpc/framework/slb"
)

type FailFastClusterInvoker struct {
	Directory
	Next    framework.Invoker
	LB      slb.LoadBalance
	Retries int32
}

func NewFailFastClusterInvoker(invoker framework.Invoker, lb slb.LoadBalance, retris int32) *FailFastClusterInvoker {
	return &FailFastClusterInvoker{
		Directory: GetDirectory(),
		Next:      invoker,
		LB:        lb,
		Retries:   retris,
	}
}

func (f *FailFastClusterInvoker) Invoke(context *framework.InvokerContext) (*entity.Result, error) {
	var err error
	//do sth

	//directory
	urls := f.Directory.ListProviderUrls(context.UniqueMetaName)
	//loadbalance
	selectedUrl := f.selectProvider(urls, context.Invocation)
	//lookup provider
	provider, err := f.Directory.LookupProvider(context.UniqueMetaName, selectedUrl)
	context.Provider = provider
	//call next
	result, err := f.Next.Invoke(context)
	if err == nil {
		return result, err
	}
	return nil, err
}

func (f *FailFastClusterInvoker) selectProvider(urls []*config.URL, invocation *entity.Invocation) *config.URL {
	if len(urls) == 1 {
		return urls[0]
	}
	return slb.GetLoadBalance("").Select(urls, invocation)
}
