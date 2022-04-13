package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetPasswordReq struct {
	// 新密码。

	Newpassword string `json:"newpassword"`
}

func (o ResetPasswordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordReq struct{}"
	}

	return strings.Join([]string{"ResetPasswordReq", string(data)}, " ")
}
