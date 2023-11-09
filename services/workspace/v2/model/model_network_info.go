package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkInfo 桌面网络信息：vpc、子网、私有ip、弹性ip、安全组
type NetworkInfo struct {
	VpcInfo *Vpc `json:"vpc_info,omitempty"`

	SubnetInfo *Subnet `json:"subnet_info,omitempty"`

	PortInfo *Port `json:"port_info,omitempty"`

	PublicIpInfo *PublicIp `json:"public_ip_info,omitempty"`

	// 桌面绑定的安全组列表
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`
}

func (o NetworkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkInfo struct{}"
	}

	return strings.Join([]string{"NetworkInfo", string(data)}, " ")
}
