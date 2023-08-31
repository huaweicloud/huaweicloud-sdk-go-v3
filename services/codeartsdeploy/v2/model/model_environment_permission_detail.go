package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvironmentPermissionDetail 环境权限详情
type EnvironmentPermissionDetail struct {

	// 是否有删除环境权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有部署权限
	CanDeploy *bool `json:"can_deploy,omitempty"`

	// 是否有编辑环境权限
	CanEdit *bool `json:"can_edit,omitempty"`

	// 是否有编辑环境权限矩阵的权限
	CanManage *bool `json:"can_manage,omitempty"`

	// 是否有环境的查看权限
	CanView *bool `json:"can_view,omitempty"`
}

func (o EnvironmentPermissionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentPermissionDetail struct{}"
	}

	return strings.Join([]string{"EnvironmentPermissionDetail", string(data)}, " ")
}
