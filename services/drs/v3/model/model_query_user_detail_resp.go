package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迁移用户响应体
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

	Privileges *[]string `json:"privileges,omitempty"`
	// 密码。

	Password *string `json:"password,omitempty"`
	// 账号拥有的角色

	Roles *[]string `json:"roles,omitempty"`
	// 是否选择。

	Selected *bool `json:"selected,omitempty"`
}

func (o QueryUserDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryUserDetailResp struct{}"
	}

	return strings.Join([]string{"QueryUserDetailResp", string(data)}, " ")
}
