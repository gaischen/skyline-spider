package rpc_client

import (
	"context"
	"github.com/skyline/skyline-spider/rpc/framework/config"
	"github.com/skyline/skyline-spider/rpc/transport"
	"sync"
)

type Directory interface {
	ListProviderUrls(serviceName string) []*config.URL
	LookupProvider(serviceName string, url *config.URL) (*transport.Provider, error)
	InvalidProvider(serviceNmae string, url *config.URL) error
	DeleteProvider(serviceNmae string, url *config.URL) error
	destroy()
}

var once sync.Once
var dr *providerDirectory

func GetDirectory() Directory {
	once.Do(func() {
		if dr == nil {
			ctx, cancelFunc := context.WithCancel(context.Background())
			dr = &providerDirectory{
				ctx:                ctx,
				cancel:             cancelFunc,
				serviceProviderMap: new(sync.Map),
				urlProviderMap:     new(sync.Map),
				urlServiceMapping:  new(sync.Map),
				providerAddLock:    sync.Mutex{},
			}
			//registerEvent(dr)
		}
	})
	return dr
}

type providerDirectory struct {
	ctx    context.Context
	cancel context.CancelFunc
	providerAddLock sync.Mutex
	serviceProviderMap *sync.Map //serviceName -> url->Provider
	urlProviderMap *sync.Map // url -> Provider
	urlServiceMapping *sync.Map //user -> serviceName
	//addressService registry.AddressService
}

func (p *providerDirectory) ListProviderUrls(serviceName string) []*config.URL {
	panic("implement me")
}

func (p *providerDirectory) LookupProvider(serviceName string, url *config.URL) (*transport.Provider, error) {
	panic("implement me")
}

func (p *providerDirectory) InvalidProvider(serviceNmae string, url *config.URL) error {
	panic("implement me")
}

func (p *providerDirectory) DeleteProvider(serviceNmae string, url *config.URL) error {
	panic("implement me")
}

func (p *providerDirectory) destroy() {
	panic("implement me")
}
