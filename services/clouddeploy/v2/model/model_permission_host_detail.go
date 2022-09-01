package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主机相关权限详情类
type PermissionHostDetail struct {

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty" xml:"can_view"`

	// 是否有编辑权限
	CanEdit *bool `json:"can_edit,omitempty" xml:"can_edit"`

	// 是否有删除权限
	CanDelete *bool `json:"can_delete,omitempty" xml:"can_delete"`

	// 是否有添加主机权限
	CanAddHost *bool `json:"can_add_host,omitempty" xml:"can_add_host"`

	// 是否测试主机连通性权限
	CanConnectionTest *bool `json:"can_connection_test,omitempty" xml:"can_connection_test"`
}

func (o PermissionHostDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionHostDetail struct{}"
	}

	return strings.Join([]string{"PermissionHostDetail", string(data)}, " ")
}
