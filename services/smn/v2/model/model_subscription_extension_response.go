package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionExtensionResponse struct {

	// 是否启用验证码发送确认短信，默认为false。
	VerificationCodeEnabled *bool `json:"verification_code_enabled,omitempty"`
}

func (o SubscriptionExtensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionExtensionResponse struct{}"
	}

	return strings.Join([]string{"SubscriptionExtensionResponse", string(data)}, " ")
}
