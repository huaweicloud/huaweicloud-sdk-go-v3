package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 一键重置裸金属服务器密码接口请求结构体
type ResetPasswordBody struct {
	ResetPassword *ResetPassword `json:"reset-password"`
}

func (o ResetPasswordBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordBody struct{}"
	}

	return strings.Join([]string{"ResetPasswordBody", string(data)}, " ")
}
