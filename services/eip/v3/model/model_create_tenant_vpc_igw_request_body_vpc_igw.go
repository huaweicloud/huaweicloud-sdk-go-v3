package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantVpcIgwRequestBodyVpcIgw 创建虚拟IGW的请求体
type CreateTenantVpcIgwRequestBodyVpcIgw struct {

	// vpcid
	VpcId string `json:"vpc_id"`

	// 创建VPC IGW的network id
	NetworkId *string `json:"network_id,omitempty"`

	// 是否添加默认路由
	AddRoute *bool `json:"add_route,omitempty"`

	// 是否使能ipv6
	EnableIpv6 *bool `json:"enable_ipv6,omitempty"`

	// 虚拟IGW的名称
	Name *string `json:"name,omitempty"`
}

func (o CreateTenantVpcIgwRequestBodyVpcIgw) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantVpcIgwRequestBodyVpcIgw struct{}"
	}

	return strings.Join([]string{"CreateTenantVpcIgwRequestBodyVpcIgw", string(data)}, " ")
}
