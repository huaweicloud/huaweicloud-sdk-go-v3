package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultProjectPermissionResponseBodyResult 权限信息
type ShowDefaultProjectPermissionResponseBodyResult struct {

	// 角色ID
	RoleId *int32 `json:"role_id,omitempty"`

	// 角色名
	RoleName *string `json:"role_name,omitempty"`

	// 是否有复制任务权限
	IsCopy *bool `json:"is_copy,omitempty"`

	// 是否有删除任务权限
	IsDelete *bool `json:"is_delete,omitempty"`

	// 是否有执行任务权限
	IsExecute *bool `json:"is_execute,omitempty"`

	// 是否有禁用任务权限
	IsForbidden *bool `json:"is_forbidden,omitempty"`

	// 是否有管理任务权限
	IsManager *bool `json:"is_manager,omitempty"`

	// 是否有修改任务权限
	IsModify *bool `json:"is_modify,omitempty"`

	// 是否有查看任务权限
	IsView *bool `json:"is_view,omitempty"`
}

func (o ShowDefaultProjectPermissionResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultProjectPermissionResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ShowDefaultProjectPermissionResponseBodyResult", string(data)}, " ")
}
