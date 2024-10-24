package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutCustomRoleToPermissionSetReqBody the request body of PutCustomRoleToPermissionSet
type PutCustomRoleToPermissionSetReqBody struct {

	// 要附加到权限集的自定义策略
	CustomRole string `json:"custom_role"`
}

func (o PutCustomRoleToPermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCustomRoleToPermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"PutCustomRoleToPermissionSetReqBody", string(data)}, " ")
}
