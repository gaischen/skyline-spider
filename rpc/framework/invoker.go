package framework

import "github.com/skyline/skyline-spider/rpc/framework/entity"

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

type InvokerContext struct {

}
