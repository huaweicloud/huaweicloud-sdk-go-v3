package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantVpcIgwRequestBodyVpcIgw 修改虚拟IGW的请求体
type UpdateTenantVpcIgwRequestBodyVpcIgw struct {

	// 虚拟IGW的名称
	Name *string `json:"name,omitempty"`

	// 是否使能ipv6
	EnableIpv6 *bool `json:"enable_ipv6,omitempty"`
}

func (o UpdateTenantVpcIgwRequestBodyVpcIgw) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantVpcIgwRequestBodyVpcIgw struct{}"
	}

	return strings.Join([]string{"UpdateTenantVpcIgwRequestBodyVpcIgw", string(data)}, " ")
}
