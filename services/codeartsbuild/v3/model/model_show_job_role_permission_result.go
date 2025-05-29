package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowJobRolePermissionResult struct {

	// 数据库中ID
	Id *int32 `json:"id,omitempty"`

	// 角色ID
	RoleId *int32 `json:"role_id,omitempty"`

	// 角色名
	RoleName *string `json:"role_name,omitempty"`

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 是否有修改任务权限
	IsModify *bool `json:"is_modify,omitempty"`

	// 是否有删除任务权限
	IsDelete *bool `json:"is_delete,omitempty"`

	// 是否有查看任务权限
	IsView *bool `json:"is_view,omitempty"`

	// 是否有执行任务权限
	IsExecute *bool `json:"is_execute,omitempty"`

	// 是否有复制任务权限
	IsCopy *bool `json:"is_copy,omitempty"`

	// 是否有禁用任务权限
	IsForbidden *bool `json:"is_forbidden,omitempty"`

	// 是否有管理任务权限
	IsManager *bool `json:"is_manager,omitempty"`

	// 任务创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 任务最后修改时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 次数
	Count *int32 `json:"count,omitempty"`
}

func (o ShowJobRolePermissionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRolePermissionResult struct{}"
	}

	return strings.Join([]string{"ShowJobRolePermissionResult", string(data)}, " ")
}
