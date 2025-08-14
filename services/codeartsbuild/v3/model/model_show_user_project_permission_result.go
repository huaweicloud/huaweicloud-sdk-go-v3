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

	// 用户是否有删除权限。
	DeletePermission *bool `json:"delete_permission,omitempty"`

	// 用户是否有查看权限。
	ViewPermission *bool `json:"view_permission,omitempty"`

	// 用户是否有执行权限。
	ExecutePermission *bool `json:"execute_permission,omitempty"`

	// 用户是否有克隆权限。
	CopyPermission *bool `json:"copy_permission,omitempty"`

	// 用户是否有禁用权限。
	ForbiddenPermission *bool `json:"forbidden_permission,omitempty"`

	// 用户是否有管理权限。
	ManagerPermission *bool `json:"manager_permission,omitempty"`

	// 角色ID
	RoleId *string `json:"role_id,omitempty"`

	// 角色名
	RoleName *string `json:"role_name,omitempty"`

	// 角色编码集合。
	RoleIds *[]string `json:"role_ids,omitempty"`

	// 角色名称集合。
	RoleNames *[]string `json:"role_names,omitempty"`
}

func (o ShowUserProjectPermissionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserProjectPermissionResult struct{}"
	}

	return strings.Join([]string{"ShowUserProjectPermissionResult", string(data)}, " ")
}
