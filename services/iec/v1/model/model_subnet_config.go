package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetConfig 子网配置
type SubnetConfig struct {

	// 边缘子网ID。
	Id string `json:"id"`

	// 创建实例是否开启IPv6能力。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 是否使用IPv6带宽。  约束： - ipv6_enable=true时，有效； - 使用IPv6带宽后，优先根据ipv6_bandwidth中配置的带宽，如果ipv6_bandwidth未设置，则使用使用IPv6子网所在Ipv6池的带宽,如果当前IPv6所在池子下面没有带宽，则自动创建带宽
	Ipv6BandwidthEnable *bool `json:"ipv6_bandwidth_enable,omitempty"`

	Ipv6Bandwidth *Ipv6Bandwidth `json:"ipv6_bandwidth,omitempty"`

	// - 功能说明：IP/Mac对列表 - 约束：     IP地址不允许为 “0.0.0.0/0”     如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。     如果allowed_address_pairs为“1.1.1.1/0”，表示关闭源目地址检查开关
	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`
}

func (o SubnetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetConfig struct{}"
	}

	return strings.Join([]string{"SubnetConfig", string(data)}, " ")
}
