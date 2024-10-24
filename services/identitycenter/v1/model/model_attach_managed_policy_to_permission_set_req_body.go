package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachManagedPolicyToPermissionSetReqBody 请求体
type AttachManagedPolicyToPermissionSetReqBody struct {

	// IAM系统身份策略唯一标识
	ManagedPolicyId string `json:"managed_policy_id"`

	// IAM系统身份策略名称
	ManagedPolicyName *string `json:"managed_policy_name,omitempty"`
}

func (o AttachManagedPolicyToPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachManagedPolicyToPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"AttachManagedPolicyToPermissionSetReqBody", string(data)}, " ")
}
