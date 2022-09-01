package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迁移用户响应体
type QueryUserDetailResp struct {

	// 用户账户id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 账户。
	Account *string `json:"account,omitempty" xml:"account"`

	// 说明。
	Comment *string `json:"comment,omitempty" xml:"comment"`

	// 是否支持迁移
	IsTransfer *bool `json:"is_transfer,omitempty" xml:"is_transfer"`

	// 权限
	Privileges *[]string `json:"privileges,omitempty" xml:"privileges"`

	// 密码。
	Password *string `json:"password,omitempty" xml:"password"`

	// 账号拥有的角色
	Roles *[]string `json:"roles,omitempty" xml:"roles"`

	// 是否选择。
	Selected *bool `json:"selected,omitempty" xml:"selected"`

	// 无法同步的用户权限
	NoPrivileges *string `json:"no_privileges,omitempty" xml:"no_privileges"`

	// 父用户
	ParentAccount *string `json:"parent_account,omitempty" xml:"parent_account"`

	// 无法同步父子关系的父用户
	NoParentAccount *string `json:"no_parent_account,omitempty" xml:"no_parent_account"`
}

func (o QueryUserDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryUserDetailResp struct{}"
	}

	return strings.Join([]string{"QueryUserDetailResp", string(data)}, " ")
}
