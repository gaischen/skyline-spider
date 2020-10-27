package rpc_client

import (
	"context"
	"github.com/juju/errors"
	"github.com/skyline/skyline-spider/rpc/framework"
	"github.com/skyline/skyline-spider/rpc/framework/config"
	"github.com/skyline/skyline-spider/rpc/framework/registry"
	"github.com/skyline/skyline-spider/rpc/framework/slb"
	"strconv"
	"time"
)

var rpcConsumerConfigMap = make(map[string]*rpcConsumerWrapper)
var registryContext, registryCancelFunc = context.WithCancel(context.Background())

//most important
func ServiceLoad() error {
	rpcConfig := config.GetRpcConfig()
	err := loadRpcConsumer(rpcConfig)

	return err
}

//load consumer proxy
func loadRpcConsumer(rpcConfig *config.RpcConfig) error {
	for _, consumerWrapper := range rpcConsumerConfigMap {
		err := LoadService(consumerWrapper, rpcConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadService(wrapper *rpcConsumerWrapper, rpcConfig *config.RpcConfig) error {
	serviceMeta := newServiceMetadataReferConfig(wrapper.referConfig)
	ctx, cancel := context.WithCancel(registryContext)
	ch, err := registry.GetMetadataService().Subscribe(ctx, serviceMeta)
	if err != nil {
		return err
	}
	if rpcConfig.ReferCheck {
		t := time.NewTimer(time.Duration(rpcConfig.RegistryConfig.SubscribeTimeout * int64(time.Millisecond)))
		select {
		case <-t.C:
			cancel()
			return errors.Annotatef(errors.Errorf(""), "")
		case <-ch:
		}
	}

	chain := framework.NewInvokerChain()

	//add load balance
	chain.AddInvoker(func(invoker framework.Invoker) framework.Invoker {
		//do sth
		lb:=slb.GetLoadBalance("")
		return NewFailFastClusterInvoker(invoker, lb, wrapper.referConfig.Retries)
	})

	providerInvoker := NewProviderInvoker()

	return nil
}

func newServiceMetadataReferConfig(referConfig *ReferConfig) *registry.ServiceMetadata {
	meta := registry.NewServiceMetadata(referConfig.JavaClassName, referConfig.Version, referConfig.Group)
	if referConfig.HashNodes != 0 {
		meta.ServiceProps[framework.HashNodes] = strconv.Itoa(referConfig.HashNodes)
	}

	if referConfig.HashArgsIndex != "" {
		meta.ServiceProps[framework.HashArgsIndex] = referConfig.HashArgsIndex
	}

	if referConfig.MethodSpecials != nil {
		for _, m := range referConfig.MethodSpecials {
			if m.HashArgsIndex != "" {
				meta.ServiceProps[m.MethodName+"."+framework.HashArgsIndex] = m.HashArgsIndex
			}
		}
	}
	return meta
}
