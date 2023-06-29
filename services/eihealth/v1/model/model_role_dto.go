package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleDto 角色
type RoleDto struct {

	// 角色id
	Id *string `json:"id,omitempty"`

	// 角色名
	Name *string `json:"name,omitempty"`
}

func (o RoleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleDto struct{}"
	}

	return strings.Join([]string{"RoleDto", string(data)}, " ")
}
