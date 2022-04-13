package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SendSmsVerificationCodeRequest struct {
	Body *SendSmVerificationCodeReq `json:"body,omitempty"`
}

func (o SendSmsVerificationCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendSmsVerificationCodeRequest struct{}"
	}

	return strings.Join([]string{"SendSmsVerificationCodeRequest", string(data)}, " ")
}
