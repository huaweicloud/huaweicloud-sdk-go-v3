package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvironmentHostPermission 环境下主机权限详情
type EnvironmentHostPermission struct {

	// 是否有删除权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有部署权限
	CanDeploy *bool `json:"can_deploy,omitempty"`

	// 是否有编辑权限
	CanEdit *bool `json:"can_edit,omitempty"`

	// 是否有权限管理权限
	CanManage *bool `json:"can_manage,omitempty"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`
}

func (o EnvironmentHostPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentHostPermission struct{}"
	}

	return strings.Join([]string{"EnvironmentHostPermission", string(data)}, " ")
}
