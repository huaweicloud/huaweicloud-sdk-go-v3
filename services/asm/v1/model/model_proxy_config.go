package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProxyConfig 网格数据面代理配置。
type ProxyConfig struct {

	// 拦截对外访问的地址范围，以半角英文逗号（,）分隔的IP网段
	IncludeIPRanges *string `json:"includeIPRanges,omitempty"`

	// 排除拦截对外访问的地址范围，以半角英文逗号（,）分隔的IP网段
	ExcludeIPRanges *string `json:"excludeIPRanges,omitempty"`

	// 排除拦截对外访问端口，以半角英文逗号（,）分隔的出站端口列表
	ExcludeOutboundPorts *string `json:"excludeOutboundPorts,omitempty"`

	// 排除拦截访问服务的端口，以半角英文逗号（,）分隔的入站端口列表
	ExcludeInboundPorts *string `json:"excludeInboundPorts,omitempty"`

	// 拦截对外访问端口，以半角英文逗号（,）分隔的出站端口列表
	IncludeOutboundPorts *string `json:"includeOutboundPorts,omitempty"`

	// 拦截访问服务的端口，以半角英文逗号（,）分隔的入站端口列表
	IncludeInboundPorts *string `json:"includeInboundPorts,omitempty"`
}

func (o ProxyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyConfig struct{}"
	}

	return strings.Join([]string{"ProxyConfig", string(data)}, " ")
}
