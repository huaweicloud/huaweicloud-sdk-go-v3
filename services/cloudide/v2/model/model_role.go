package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Role struct {
	// 子角色

	CRole *string `json:"c_role,omitempty"`
	// id

	Id *string `json:"id,omitempty"`
	// 角色

	Role *string `json:"role,omitempty"`
	// 角色执行操作列表

	RoleActionses *[]RoleAction `json:"role_actionses,omitempty"`
}

func (o Role) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Role struct{}"
	}

	return strings.Join([]string{"Role", string(data)}, " ")
}
