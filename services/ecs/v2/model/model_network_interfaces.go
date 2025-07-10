package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkInterfaces struct {

	// 网卡端口id
	PortId *string `json:"port_id,omitempty"`

	// 是否是主网卡
	Primary *bool `json:"primary,omitempty"`

	// ipv4地址
	IpAddresses *[]string `json:"ip_addresses,omitempty"`

	// ipv6地址
	Ipv6Addresses *[]string `json:"ipv6_addresses,omitempty"`

	Association *Association `json:"association,omitempty"`
}

func (o NetworkInterfaces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkInterfaces struct{}"
	}

	return strings.Join([]string{"NetworkInterfaces", string(data)}, " ")
}
