package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionHostDetailNew 主机相关权限详情类
type PermissionHostDetailNew struct {

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有编辑权限
	CanEdit *bool `json:"can_edit,omitempty"`

	// 是否有删除权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有添加主机权限
	CanAddHost *bool `json:"can_add_host,omitempty"`

	// 是否有复制主机权限
	CanCopy *bool `json:"can_copy,omitempty"`
}

func (o PermissionHostDetailNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionHostDetailNew struct{}"
	}

	return strings.Join([]string{"PermissionHostDetailNew", string(data)}, " ")
}
