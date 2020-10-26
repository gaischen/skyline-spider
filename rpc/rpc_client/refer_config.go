package rpc_client

import (
	"github.com/juju/errors"
	"github.com/skyline/skyline-spider/rpc/framework"
	"github.com/skyline/skyline-spider/rpc/framework/serializations"
	"strings"
	"time"
)

type MethodSpecial struct {
	MethodName    string
	Async         bool
	OneWay        bool
	Timeout       int
	LoadBalance   string
	HashArgsIndex string
}

type ReferConfig struct {
	//是否泛化
	Generic    bool
	Retries    int32
	SerialType serializations.SERIALIZATION

	Fallback      map[string]interface{}
	JavaClassName string
	Version       string
	Group         string
	Timeout       time.Duration
	LoadBalance   string

	HashNodes     int
	HashArgsIndex string

	UniqueMetaName    string
	UniqueServiceName string

	//ipv4:port
	TargetAddress []string

	MethodSpecials []*MethodSpecial
	IsInit         bool

	genericService *GenericService
}

func NewReferConfig(javaClass string, timeout time.Duration, targetAddress string) *ReferConfig {
	if "" == javaClass {
		panic("java class is null")
	}
	r := &ReferConfig{}
	r.JavaClassName = javaClass
	r.SetTargetAddresses(targetAddress)
	r.Timeout = timeout
	r.Retries = framework.DEFAULT_RETRIES_TIMES
	r.Group = framework.DEFAULT_GRUOP
	return r
}

func (r *ReferConfig) Init() {
	if r.IsInit {
		return
	}

	if r.JavaClassName == "" {
		panic("target class is null")
	}

	sb := strings.Builder{}
	sb.WriteString(r.JavaClassName)

	if r.Version != "" {
		sb.WriteString(":")
		sb.WriteString(r.Version)
	}
	r.UniqueServiceName = sb.String()

	if r.Group != "" {
		sb.WriteString("@")
		sb.WriteString(r.Group)
	}
	r.UniqueMetaName = sb.String()

	r.IsInit = true
}

func (r *ReferConfig) SetTargetAddresses(targetAddresses string) {
	if targetAddresses == "" {
		return
	}

	addresses := strings.Split(targetAddresses, ",")
	r.TargetAddress = make([]string, len(addresses))
	for i, a := range addresses {
		r.TargetAddress[i] = a
	}
}

func (r *ReferConfig) AddMethodSpecial(special *MethodSpecial) {
	r.MethodSpecials = append(r.MethodSpecials, special)
}

func (r *ReferConfig) SetGeneric(generic bool) {
	r.Generic = generic
}

func (r *ReferConfig) GetGenericService() (*GenericService, error) {

	if !r.Generic {
		return nil, errors.Errorf("service %s not support generic call", r.UniqueServiceName)
	}
	if r.genericService == nil {
		return nil, errors.New("generic service is nil,please load referConfig before call generic method")
	}
	return r.genericService, nil
}
