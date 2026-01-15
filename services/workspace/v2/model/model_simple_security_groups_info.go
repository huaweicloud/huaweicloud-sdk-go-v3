package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleSecurityGroupsInfo 安全组。
type SimpleSecurityGroupsInfo struct {

	// 安全组描述。
	Description *string `json:"description,omitempty"`

	// 安全组名称。
	Name *string `json:"name,omitempty"`

	// 安全组的主键。
	Id *string `json:"id,omitempty"`

	// 外部网关信息。
	SecurityGroupRules *string `json:"security_group_rules,omitempty"`

	// 租户主键。
	TenantId *string `json:"tenant_id,omitempty"`

	// 资源租户主键。
	ProjectId *string `json:"project_id,omitempty"`

	// 资源创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 资源更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o SimpleSecurityGroupsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleSecurityGroupsInfo struct{}"
	}

	return strings.Join([]string{"SimpleSecurityGroupsInfo", string(data)}, " ")
}
