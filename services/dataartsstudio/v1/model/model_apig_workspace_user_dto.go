package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigWorkspaceUserDto 获取DataArtsStudio工作空间角色信息
type ApigWorkspaceUserDto struct {

	// 用户类型，0:添加用户;1:添加用户组
	Type int32 `json:"type"`

	// 用户列表信息
	UserIds *[]ApigIamUserDto `json:"user_ids,omitempty"`

	// 用户组列表信息
	Groups *[]Group `json:"groups,omitempty"`

	// 空间角色列表
	RolesIds []ApigRole `json:"roles_ids"`
}

func (o ApigWorkspaceUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigWorkspaceUserDto struct{}"
	}

	return strings.Join([]string{"ApigWorkspaceUserDto", string(data)}, " ")
}
