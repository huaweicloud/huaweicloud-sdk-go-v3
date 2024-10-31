package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetInfo 管理子网信息
type SubnetInfo struct {

	// 子网所在的可用区标识，从终端节点获取，参考[终端节点](cfw_02_0003.xml)
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 功能说明：虚拟私有云下可用子网的范围 取值范围： 10.0.0.0/8~24 172.16.0.0/12~24 192.168.0.0/16~24 不指定cidr时，默认值为空 约束：必须是cidr格式，例如:192.168.0.0/16
	Cidr *string `json:"cidr,omitempty"`

	// 子网名称
	Name *string `json:"name,omitempty"`

	// 子网id
	Id *string `json:"id,omitempty"`

	// 子网的网关，取值范围为子网网段cidr中的ip地址
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// 创建vpc产生的uuid
	VpcId *string `json:"vpc_id,omitempty"`

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
