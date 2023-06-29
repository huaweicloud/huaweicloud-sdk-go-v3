package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleAuthority 角色权限信息
type RoleAuthority struct {

	// 角色名称
	Role *string `json:"role,omitempty"`

	// 权限列表
	RightList *[]string `json:"right_list,omitempty"`
}

func (o RoleAuthority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleAuthority struct{}"
	}

	return strings.Join([]string{"RoleAuthority", string(data)}, " ")
}
