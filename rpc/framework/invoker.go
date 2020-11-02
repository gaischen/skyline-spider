package framework

import (
	"github.com/vanga-top/skyline-spider/rpc/framework/entity"
	"github.com/vanga-top/skyline-spider/rpc/transport"
)

type Invoker interface {
	Invoke(context *InvokerContext) (*entity.Result, error)
}

type InvokerFun func(invocation *InvokerContext) (*entity.Result, error)

func (i InvokerFun) Invoke(context *InvokerContext) (*entity.Result, error) {
	return i(context)
}

type Middleware func(invoker Invoker) Invoker
type InvokerChain struct {
	invokers []Middleware
}

func NewInvokerChain() *InvokerChain {
	return &InvokerChain{invokers: []Middleware{}}
}

func (ic *InvokerChain) AddInvoker(m Middleware) {
	ic.invokers = append(ic.invokers, m)
}

func (ic *InvokerChain) BuildInvokerChain(endInvoker Invoker) (Invoker, error) {
	//normally last invoker is endInvoker
	var invoker = endInvoker
	for i := len(ic.invokers) - 1; i >= 0; i-- {
		invoker = ic.invokers[i](invoker)
	}
	return invoker, nil
}

type ResponseCallback interface {
	Callback(response *CallBackResponse)
}
type CallBackResponse struct {
	Response interface{}
	Err      error
}

type AsyncListener func(response *CallBackResponse, context *InvokerContext)

type InvokerContext struct {
	UniqueMetaName string
	Invocation     *entity.Invocation
	Provider       *transport.Provider
	Callback       transport.Callback
	context        map[string]interface{}
	asyncListeners []AsyncListener
}
