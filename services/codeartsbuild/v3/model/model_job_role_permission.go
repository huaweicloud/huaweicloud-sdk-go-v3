package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobRolePermission struct {

	// 主键ID
	Id *int32 `json:"id,omitempty"`

	// 角色ID
	RoleId *int32 `json:"role_id,omitempty"`

	// devuc的角色ID
	DevucRoleId *string `json:"devuc_role_id,omitempty"`

	// 角色名称
	RoleName *string `json:"role_name,omitempty"`

	// 修改权限
	IsModify *bool `json:"is_modify,omitempty"`

	// 删除权限
	IsDelete *bool `json:"is_delete,omitempty"`

	// 查看权限
	IsView *bool `json:"is_view,omitempty"`

	// 执行权限
	IsExecute *bool `json:"is_execute,omitempty"`

	// 复制权限
	IsCopy *bool `json:"is_copy,omitempty"`

	// 禁用权限
	IsForbidden *bool `json:"is_forbidden,omitempty"`

	// 管理权限
	IsManager *bool `json:"is_manager,omitempty"`

	// 数量
	Count *int32 `json:"count,omitempty"`
}

func (o JobRolePermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRolePermission struct{}"
	}

	return strings.Join([]string{"JobRolePermission", string(data)}, " ")
}
