package registry

import "strings"

type ServiceMetadata struct {
	Name string
	Version string
	Group string
	UniqueServiceName string
	UniqueMetaName string
	ServiceProps map[string]string
}

func NewServiceMetadata(name string, version string, group string) *ServiceMetadata {
	meta := &ServiceMetadata{}
	meta.Name = name

	if version != "" {
		meta.Version = version
	}

	if group != "" {
		meta.Group = group
	}

	meta.ServiceProps = make(map[string]string)

	meta.init()
	return meta
}

func (m *ServiceMetadata) init() {

	sb := strings.Builder{}
	sb.WriteString(m.Name)
	if m.Version != "" {
		sb.WriteString(":")
		sb.WriteString(m.Version)
	}
	m.UniqueServiceName = sb.String()

	if m.Group != "" {
		sb.WriteString("@")
		sb.WriteString(m.Group)
	}

	m.UniqueMetaName = sb.String()
}

func (m *ServiceMetadata) Properties() map[string]string {
	return m.ServiceProps
}

func (m *ServiceMetadata) AddProperty(key, value string) {
	m.ServiceProps[key] = value
}

func (m *ServiceMetadata) GetProperty(key string) string {
	return m.ServiceProps[key]
}
