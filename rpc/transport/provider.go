package transport

import (
	"github.com/vanga-top/skyline-spider/rpc/framework/entity"
	"github.com/vanga-top/skyline-spider/rpc/framework/serializations"
	"time"
)

type Provider struct {

}

type NetCall struct {
	Done  chan struct{}
	Error error

	SerType         serializations.SERIALIZATION
	Timeout         time.Duration
	ExecuteTime     time.Duration //执行时间 oneway情况下为空
	SerializeTime   time.Duration
	DeSerializeTime time.Duration
	SerializeSize   uint32
	DeSerializeSize uint32

	Invocation *entity.Invocation
	Response   interface{}
}

type Callback func(call *NetCall)