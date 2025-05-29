package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserProjectPermissionResult 结果
type ShowUserProjectPermissionResult struct {

	// 工程编号
	ProjectId *string `json:"project_id,omitempty"`

	// 用户是否有创建权限
	CreatePermission *bool `json:"create_permission,omitempty"`

	// 用户是否有修改权限
	ModifyPermission *bool `json:"modify_permission,omitempty"`

	// 用户是否有分类权限
	GroupPermission *bool `json:"group_permission,omitempty"`

	// 角色ID
	RoleId *string `json:"role_id,omitempty"`

	// 角色名
	RoleName *string `json:"role_name,omitempty"`
}

func (o ShowUserProjectPermissionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserProjectPermissionResult struct{}"
	}

	return strings.Join([]string{"ShowUserProjectPermissionResult", string(data)}, " ")
}
