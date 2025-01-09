package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InternetInfo struct {

	// VPC名称
	VpcName *string `json:"vpc_name,omitempty"`

	// 子网id
	SubnetId *string `json:"subnet_id,omitempty"`

	// 子网名称
	SubnetName *string `json:"subnet_name,omitempty"`

	// NAT id
	NatId *string `json:"nat_id,omitempty"`

	// NAT名称
	NatName *string `json:"nat_name,omitempty"`

	// 弹性公网IP
	Eip *string `json:"eip,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o InternetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InternetInfo struct{}"
	}

	return strings.Join([]string{"InternetInfo", string(data)}, " ")
}
