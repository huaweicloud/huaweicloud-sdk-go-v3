package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryUserDetailResp 迁移用户响应体
type QueryUserDetailResp struct {

	// 用户账户id。
	Id *string `json:"id,omitempty"`

	// 账户。
	Account *string `json:"account,omitempty"`

	// 说明。
	Comment *string `json:"comment,omitempty"`

	// 是否支持迁移
	IsTransfer *bool `json:"is_transfer,omitempty"`

	// 权限
	Privileges *string `json:"privileges,omitempty"`

	// 密码。
	Password *string `json:"password,omitempty"`

	// 账号拥有的角色
	Roles *[]string `json:"roles,omitempty"`

	// 是否选择。
	Selected *bool `json:"selected,omitempty"`

	// 无法同步的用户权限
	NoPrivileges *string `json:"no_privileges,omitempty"`

	// 父用户
	ParentAccount *string `json:"parent_account,omitempty"`

	// 无法同步父子关系的父用户
	NoParentAccount *string `json:"no_parent_account,omitempty"`
}

func (o QueryUserDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryUserDetailResp struct{}"
	}

	return strings.Join([]string{"QueryUserDetailResp", string(data)}, " ")
}
