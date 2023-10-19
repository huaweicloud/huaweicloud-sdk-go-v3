package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetInfo 管理子网信息
type SubnetInfo struct {

	// 子网id
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// vpc cidr
	Cidr *string `json:"cidr,omitempty"`

	// 子网名称
	Name *string `json:"name,omitempty"`

	// 子网id
	Id *string `json:"id,omitempty"`

	// 子网网关ip
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// vpc id
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网的状态
	Status *string `json:"status,omitempty"`

	// 是否支持ipv6，boolean值为true表示是，false表示否
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
}

func (o SubnetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetInfo struct{}"
	}

	return strings.Join([]string{"SubnetInfo", string(data)}, " ")
}
