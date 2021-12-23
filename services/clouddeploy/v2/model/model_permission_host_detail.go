package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主机相关权限详情类
type PermissionHostDetail struct {
	// 是否有查看权限

	CanView *bool `json:"can_view,omitempty"`
	// 是否有编辑权限

	CanEdit *bool `json:"can_edit,omitempty"`
	// 是否有删除权限

	CanDelete *bool `json:"can_delete,omitempty"`
	// 是否有添加主机权限

	CanAddHost *bool `json:"can_add_host,omitempty"`
	// 是否测试主机连通性权限

	CanConnectionTest *bool `json:"can_connection_test,omitempty"`
}

func (o PermissionHostDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionHostDetail struct{}"
	}

	return strings.Join([]string{"PermissionHostDetail", string(data)}, " ")
}
