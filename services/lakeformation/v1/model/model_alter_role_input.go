package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlterRoleInput role信息
type AlterRoleInput struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 角色名称
	RoleName *string `json:"role_name,omitempty"`
}

func (o AlterRoleInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterRoleInput struct{}"
	}

	return strings.Join([]string{"AlterRoleInput", string(data)}, " ")
}
