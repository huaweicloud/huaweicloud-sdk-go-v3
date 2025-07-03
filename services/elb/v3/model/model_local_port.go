package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocalPort
type LocalPort struct {

	// port ID。
	PortId *string `json:"port_id,omitempty"`

	// port IPv4地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// port IPv6地址。
	Ipv6Address *string `json:"ipv6_address,omitempty"`

	// port 类型。  取值范围： - l4_hc：四层dnat转发时健康检查使用的地址。 - l4 四层fullnat 转发时健康检查及业务转发使用的地址。 - l7 七层健康检查及业务转发使用的地址。
	Type *string `json:"type,omitempty"`

	// port 虚拟子网ID。
	VirsubnetId *string `json:"virsubnet_id,omitempty"`
}

func (o LocalPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalPort struct{}"
	}

	return strings.Join([]string{"LocalPort", string(data)}, " ")
}
