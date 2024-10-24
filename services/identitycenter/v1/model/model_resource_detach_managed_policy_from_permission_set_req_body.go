package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDetachManagedPolicyFromPermissionSetReqBody DetachManagedPolicyFromPermissionSetReqBody请求体
type ResourceDetachManagedPolicyFromPermissionSetReqBody struct {

	// IAM系统策略唯一标识
	ManagedRoleId string `json:"managed_role_id"`
}

func (o ResourceDetachManagedPolicyFromPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDetachManagedPolicyFromPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"ResourceDetachManagedPolicyFromPermissionSetReqBody", string(data)}, " ")
}
