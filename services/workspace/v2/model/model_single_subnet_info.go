package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SingleSubnetInfo struct {

	// 子网ID。
	Id *string `json:"id,omitempty"`

	// 是否未被workspace真正使用。如果未使用，则表示可以从workspace中删除。
	Unused *bool `json:"unused,omitempty"`

	// 是否已被workspace选择使用。
	Checked *bool `json:"checked,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网名称。
	Name *string `json:"name,omitempty"`

	// 子网网段。
	Cidr *string `json:"cidr,omitempty"`

	// 子网网关。
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// 是否支持DHCP。
	DhcpEnable *bool `json:"dhcp_enable,omitempty"`

	// 是否支持IPV6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 主DNS。
	PrimaryDns *string `json:"primary_dns,omitempty"`

	// 备DNS。
	SecondaryDns *string `json:"secondary_dns,omitempty"`

	// 子网状态。
	Status *string `json:"status,omitempty"`

	// 可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 网络id。
	NeutronSubnetId *string `json:"neutron_subnet_id,omitempty"`

	// ipv6网络id。
	NeutronIpv6SubnetId *string `json:"neutron_ipv6_subnet_id,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o SingleSubnetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleSubnetInfo struct{}"
	}

	return strings.Join([]string{"SingleSubnetInfo", string(data)}, " ")
}
