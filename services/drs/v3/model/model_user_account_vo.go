package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户信息
type UserAccountVo struct {

	// 用户账户ID。
	Id string `json:"id" xml:"id"`

	// 用户
	Account string `json:"account" xml:"account"`

	// 说明
	Comment *string `json:"comment,omitempty" xml:"comment"`

	// 是否支持迁移
	IsTransfer bool `json:"is_transfer" xml:"is_transfer"`

	// 权限列表
	Privileges *[]string `json:"privileges,omitempty" xml:"privileges"`

	// 密码
	Password *string `json:"password,omitempty" xml:"password"`

	// 是否重置密码。
	IsSetPassword *bool `json:"is_set_password,omitempty" xml:"is_set_password"`

	// 角色
	Roles []string `json:"roles" xml:"roles"`

	// 是否选择。
	Selected bool `json:"selected" xml:"selected"`
}

func (o UserAccountVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserAccountVo struct{}"
	}

	return strings.Join([]string{"UserAccountVo", string(data)}, " ")
}
