package rpc_client

import "time"

type RpcConsumer interface {
	ConsumerClassName() string
}
type RpcConsumerFunc func() string

func (t RpcConsumerFunc) ConsumerClassName() string {
	return t()
}

type rpcConsumerWrapper struct {
	rpcConsumer RpcConsumer
	referConfig *ReferConfig
}

func RegisterConsumer(consumer RpcConsumer) {
	referConfig := NewReferConfig(consumer.ConsumerClassName(), 3*time.Second, "")
	referConfig.Init()
	rpcConsumerConfigMap[referConfig.JavaClassName] = &rpcConsumerWrapper{
		rpcConsumer: consumer,
		referConfig: referConfig,
	}
}
