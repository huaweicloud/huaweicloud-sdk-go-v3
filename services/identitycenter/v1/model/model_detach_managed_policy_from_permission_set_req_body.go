package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachManagedPolicyFromPermissionSetReqBody DetachManagedPolicyFromPermissionSetReqBody的请求体
type DetachManagedPolicyFromPermissionSetReqBody struct {

	// 管理策略唯一标识
	ManagedPolicyId string `json:"managed_policy_id"`
}

func (o DetachManagedPolicyFromPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachManagedPolicyFromPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"DetachManagedPolicyFromPermissionSetReqBody", string(data)}, " ")
}
