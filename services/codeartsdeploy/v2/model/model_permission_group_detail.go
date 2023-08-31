package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionGroupDetail 主机集群相关权限详情类
type PermissionGroupDetail struct {

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有编辑权限
	CanEdit *bool `json:"can_edit,omitempty"`

	// 是否有删除权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有添加主机权限
	CanAddHost *bool `json:"can_add_host,omitempty"`

	// 是否有编辑主机集群权限矩阵的权限
	CanManage *bool `json:"can_manage,omitempty"`
}

func (o PermissionGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionGroupDetail struct{}"
	}

	return strings.Join([]string{"PermissionGroupDetail", string(data)}, " ")
}
