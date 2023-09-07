package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuntimePermission struct {

	// 动作列表
	Action *string `json:"action,omitempty"`

	// 动作列表
	ActionList *[]string `json:"action_list,omitempty"`

	// 权限名称
	Name *string `json:"name,omitempty"`

	// 权限名称
	PermissionName *string `json:"permission_name,omitempty"`

	// 权限类型
	PermissionType *string `json:"permission_type,omitempty"`
}

func (o RuntimePermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimePermission struct{}"
	}

	return strings.Join([]string{"RuntimePermission", string(data)}, " ")
}
