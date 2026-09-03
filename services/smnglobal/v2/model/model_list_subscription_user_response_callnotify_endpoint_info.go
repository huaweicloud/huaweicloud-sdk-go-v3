package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseCallnotifyEndpointInfo struct {

	// 终端地址。
	Endpoint string `json:"endpoint"`

	// 是否启用验证码发送确认短信，默认为false。
	VerificationCodeEnabled *bool `json:"verification_code_enabled,omitempty"`
}

func (o ListSubscriptionUserResponseCallnotifyEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseCallnotifyEndpointInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseCallnotifyEndpointInfo", string(data)}, " ")
}
