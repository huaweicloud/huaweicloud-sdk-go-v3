package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineReference struct {

	// vpc名称
	Vpc *string `json:"vpc,omitempty"`

	// 微服务引擎部署的可用区列表
	AzList *[]string `json:"azList,omitempty"`

	// 微服务引擎子网网络ID
	NetworkId *string `json:"networkId,omitempty"`

	// 微服务引擎ipv4子网划分
	SubnetCidr *string `json:"subnetCidr,omitempty"`

	// 微服务引擎ipv6子网划分
	SubnetCidrV6 *string `json:"subnetCidrV6,omitempty"`

	// 微服务引擎子网网关
	SubnetGateway *string `json:"subnetGateway,omitempty"`

	// 微服务引擎公网地址ID
	PublicIpId *string `json:"publicIpId,omitempty"`

	// 微服务引擎可支持的微服务总数
	ServiceLimit *int32 `json:"serviceLimit,omitempty"`

	// 微服务引擎可支持的实例总数
	InstanceLimit *int32 `json:"instanceLimit,omitempty"`

	// 微服务引擎附加参数
	Inputs map[string]string `json:"inputs,omitempty"`
}

func (o EngineReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineReference struct{}"
	}

	return strings.Join([]string{"EngineReference", string(data)}, " ")
}
