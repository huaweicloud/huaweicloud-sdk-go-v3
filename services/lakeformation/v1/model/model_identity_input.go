package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdentityInput 用户名称，用户组，角色三个信息
type IdentityInput struct {

	// IAM用户
	UserNames *[]string `json:"user_names,omitempty"`

	// 用户组
	Groups *[]string `json:"groups,omitempty"`

	// 角色
	Roles *[]string `json:"roles,omitempty"`
}

func (o IdentityInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityInput struct{}"
	}

	return strings.Join([]string{"IdentityInput", string(data)}, " ")
}
