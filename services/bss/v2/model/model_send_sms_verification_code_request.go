package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendSmsVerificationCodeRequest Request Object
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
