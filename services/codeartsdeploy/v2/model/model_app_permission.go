package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppPermission struct {

	// 角色id
	DevRoleId string `json:"dev_role_id"`

	// 是否具有编辑权限
	CanModify bool `json:"can_modify"`

	// 是否具有删除权限
	CanDelete bool `json:"can_delete"`

	// 是否具有查看权限
	CanView bool `json:"can_view"`

	// 是否具有执行权限
	CanExecute bool `json:"can_execute"`

	// 是否具有复制权限
	CanCopy bool `json:"can_copy"`

	// 是否具有权限管理权限
	CanManage bool `json:"can_manage"`

	// 是否具有创建环境权限
	CanCreateEnv bool `json:"can_create_env"`

	// 是否具有禁用权限
	CanDisable bool `json:"can_disable"`
}

func (o AppPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppPermission struct{}"
	}

	return strings.Join([]string{"AppPermission", string(data)}, " ")
}
