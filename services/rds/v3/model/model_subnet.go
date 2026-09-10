package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Subnet 子网信息
type Subnet struct {

	// 子网ID
	Id *string `json:"id,omitempty"`

	// 子网名字
	Name *string `json:"name,omitempty"`

	// 是否是IPV6子网
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 子网的CIDR信息
	Cidr *string `json:"cidr,omitempty"`

	// IPV6子网的CIDR信息
	CidrV6 *string `json:"cidr_v6,omitempty"`

	// 子网的网关
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// IPV6子网的网关
	GatewayIpV6 *string `json:"gateway_ip_v6,omitempty"`

	// 子网的可用区
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o Subnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subnet struct{}"
	}

	return strings.Join([]string{"Subnet", string(data)}, " ")
}
