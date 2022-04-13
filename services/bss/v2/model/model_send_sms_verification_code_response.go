package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SendSmsVerificationCodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendSmsVerificationCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendSmsVerificationCodeResponse struct{}"
	}

	return strings.Join([]string{"SendSmsVerificationCodeResponse", string(data)}, " ")
}
