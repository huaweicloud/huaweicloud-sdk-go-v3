package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetUserPasswrodReq struct {
	// 用户新密码。

	NewPassword *string `json:"new_password,omitempty"`
}

func (o ResetUserPasswrodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswrodReq struct{}"
	}

	return strings.Join([]string{"ResetUserPasswrodReq", string(data)}, " ")
}
