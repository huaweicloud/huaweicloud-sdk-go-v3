package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceAttachManagedPolicyToPermissionSetReqBody 请求体
type ResourceAttachManagedPolicyToPermissionSetReqBody struct {

	// IAM系统策略唯一标识
	ManagedRoleId string `json:"managed_role_id"`

	// IAM系统策略名称
	ManagedRoleName *string `json:"managed_role_name,omitempty"`
}

func (o ResourceAttachManagedPolicyToPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceAttachManagedPolicyToPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"ResourceAttachManagedPolicyToPermissionSetReqBody", string(data)}, " ")
}
