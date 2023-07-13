package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectRoleRsp 项目角色详情
type ProjectRoleRsp struct {

	// 项目角色名
	RoleType *string `json:"role_type,omitempty"`

	// 项目成员列表
	Users *[]BindUserRsp `json:"users,omitempty"`
}

func (o ProjectRoleRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectRoleRsp struct{}"
	}

	return strings.Join([]string{"ProjectRoleRsp", string(data)}, " ")
}
