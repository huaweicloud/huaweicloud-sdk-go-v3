package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseSmsEndpointInfo struct {

	// 终端地址。
	Endpoint string `json:"endpoint"`

	// 是否启用验证码发送确认短信，默认为false。
	VerificationCodeEnabled *bool `json:"verification_code_enabled,omitempty"`
}

func (o ListSubscriptionUserResponseSmsEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseSmsEndpointInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseSmsEndpointInfo", string(data)}, " ")
}
