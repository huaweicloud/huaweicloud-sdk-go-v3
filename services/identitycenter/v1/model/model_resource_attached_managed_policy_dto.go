package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceAttachedManagedPolicyDto 系统策略详细信息
type ResourceAttachedManagedPolicyDto struct {

	// IAM系统策略唯一标识
	RoleId *string `json:"role_id,omitempty"`

	// IAM系统策略名称
	RoleName *string `json:"role_name,omitempty"`
}

func (o ResourceAttachedManagedPolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceAttachedManagedPolicyDto struct{}"
	}

	return strings.Join([]string{"ResourceAttachedManagedPolicyDto", string(data)}, " ")
}
