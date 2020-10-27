package rpc_client

import (
	"github.com/skyline/skyline-spider/rpc/framework"
	"github.com/skyline/skyline-spider/rpc/framework/entity"
	"github.com/skyline/skyline-spider/rpc/framework/slb"
)

type FailFastClusterInvoker struct {
	Next    framework.Invoker
	LB      slb.LoadBalance
	Retries int32
}

func NewFailFastClusterInvoker(invoker framework.Invoker, lb slb.LoadBalance, retris int32) *FailFastClusterInvoker {
	return &FailFastClusterInvoker{
		Next:    invoker,
		LB:      lb,
		Retries: retris,
	}
}

func (f *FailFastClusterInvoker) Invoke(context *framework.InvokerContext) (*entity.Result, error) {
	var err error
	//do sth

	//call next
	result, err := f.Next.Invoke(context)
	if err == nil {
		return result, err
	}
	return nil, err
}
