package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachCustomerManagedPolicyToPermissionSetReqBody 请求体
type AttachCustomerManagedPolicyToPermissionSetReqBody struct {

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略内容
	PolicyContent *string `json:"policy_content,omitempty"`

	// 策略描述
	Description *string `json:"description,omitempty"`
}

func (o AttachCustomerManagedPolicyToPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachCustomerManagedPolicyToPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"AttachCustomerManagedPolicyToPermissionSetReqBody", string(data)}, " ")
}
