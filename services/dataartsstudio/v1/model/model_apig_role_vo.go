package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigRoleVo 获取DataArtsStudio工作空间角色信息
type ApigRoleVo struct {

	// 角色id
	RoleId *string `json:"role_id,omitempty"`

	// 角色编码
	RoleCode *string `json:"role_code,omitempty"`

	// 角色名称
	RoleName *string `json:"role_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o ApigRoleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigRoleVo struct{}"
	}

	return strings.Join([]string{"ApigRoleVo", string(data)}, " ")
}
