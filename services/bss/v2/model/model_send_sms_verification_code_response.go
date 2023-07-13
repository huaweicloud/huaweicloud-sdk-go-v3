package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendSmsVerificationCodeResponse Response Object
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
