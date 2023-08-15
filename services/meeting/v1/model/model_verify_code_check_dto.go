package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VerifyCodeCheckDto struct {

	// 必须和发送验证码时带的用户身份信息相同。
	User string `json:"user"`

	// 验证码。
	Code string `json:"code"`
}

func (o VerifyCodeCheckDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyCodeCheckDto struct{}"
	}

	return strings.Join([]string{"VerifyCodeCheckDto", string(data)}, " ")
}
