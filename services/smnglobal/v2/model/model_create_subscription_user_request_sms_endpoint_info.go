package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestSmsEndpointInfo struct {

	// 终端地址。必须是一个电话号码。
	Endpoint string `json:"endpoint"`

	// 是否启用验证码，默认为false。当protocol值为sms或callnotify，且该字段值设置为true时，发送订阅确认短信为验证码格式；该字段为false或者不存在时，发送的订阅确认短信为超链接格式。当protocol值为其他协议时，该字段不生效
	VerificationCodeEnabled *bool `json:"verification_code_enabled,omitempty"`
}

func (o CreateSubscriptionUserRequestSmsEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestSmsEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestSmsEndpointInfo", string(data)}, " ")
}
