package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Role struct {

	// 角色id
	Id string `json:"id"`

	// 角色名称
	DisplayName string `json:"display_name"`

	// 角色描述
	RoleDescription string `json:"role_description"`

	// 描述
	Description string `json:"description"`

	// 角色范围
	Scope string `json:"scope"`
}

func (o Role) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Role struct{}"
	}

	return strings.Join([]string{"Role", string(data)}, " ")
}
