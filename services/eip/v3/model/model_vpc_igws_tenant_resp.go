package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcIgwsTenantResp 虚拟IGW对象
type VpcIgwsTenantResp struct {

	// 虚拟IGW的uuid
	Id *string `json:"id,omitempty"`

	// 虚拟IGW的租户id
	ProjectId *string `json:"project_id,omitempty"`

	// 虚拟IGW的vpcid
	VpcId *string `json:"vpc_id,omitempty"`

	// 虚拟IGW的名称
	Name *string `json:"name,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 创建IGW使用的VPC具体子网
	NetworkId *string `json:"network_id,omitempty"`

	// 是否使能IPV6
	EnableIpv6 *bool `json:"enable_ipv6,omitempty"`
}

func (o VpcIgwsTenantResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcIgwsTenantResp struct{}"
	}

	return strings.Join([]string{"VpcIgwsTenantResp", string(data)}, " ")
}
