package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VerifyCodeCheckDto struct {
	// 必须和发送验证码时带的用户身份信息相同 maxLength：255 minLength：1

	User string `json:"user"`
	// 验证码 maxLength：32 minLength：1

	Code string `json:"code"`
}

func (o VerifyCodeCheckDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyCodeCheckDto struct{}"
	}

	return strings.Join([]string{"VerifyCodeCheckDto", string(data)}, " ")
}
