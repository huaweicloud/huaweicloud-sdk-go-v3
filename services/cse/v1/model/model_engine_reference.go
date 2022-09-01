package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineReference struct {

	// vpc名称
	Vpc *string `json:"vpc,omitempty" xml:"vpc"`

	// 微服务引擎专享版部署的可用区列表
	AzList *[]string `json:"az_list,omitempty" xml:"az_list"`

	// 微服务引擎专享版子网网络ID
	NetworkId *string `json:"network_id,omitempty" xml:"network_id"`

	// 微服务引擎专享版ipv4子网划分
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr"`

	// 微服务引擎专享版ipv6子网划分
	SubnetCidrV6 *string `json:"subnet_cidr_v6,omitempty" xml:"subnet_cidr_v6"`

	// 微服务引擎专享版子网网关
	SubnetGateway *string `json:"subnet_gateway,omitempty" xml:"subnet_gateway"`

	// 微服务引擎专享版公网地址ID
	PublicIpId *string `json:"public_ip_id,omitempty" xml:"public_ip_id"`

	// 微服务引擎专享版可支持的微服务总数
	ServiceLimit *int32 `json:"service_limit,omitempty" xml:"service_limit"`

	// 微服务引擎专享版可支持的实例总数
	InstanceLimit *int32 `json:"instance_limit,omitempty" xml:"instance_limit"`

	// 微服务引擎专享版附加参数
	Inputs map[string]string `json:"inputs,omitempty" xml:"inputs"`
}

func (o EngineReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineReference struct{}"
	}

	return strings.Join([]string{"EngineReference", string(data)}, " ")
}
