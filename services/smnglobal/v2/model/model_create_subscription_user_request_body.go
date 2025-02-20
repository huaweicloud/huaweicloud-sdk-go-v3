package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestBody struct {

	// 订阅用户名称。
	Name string `json:"name"`

	// 订阅用户分组。每个订阅分组只能包含中英文、数字([0-9])、下划线(_)，下划线不能出现在开始或结尾，下划线不能连续出现，长度为1到32个字符。
	Group *[]string `json:"group,omitempty"`

	Http *CreateSubscriptionUserRequestHttpEndpointInfo `json:"http,omitempty"`

	Https *CreateSubscriptionUserRequestHttpsEndpointInfo `json:"https,omitempty"`

	Sms *CreateSubscriptionUserRequestSmsEndpointInfo `json:"sms,omitempty"`

	Email *CreateSubscriptionUserRequestEmailEndpointInfo `json:"email,omitempty"`
}

func (o CreateSubscriptionUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestBody", string(data)}, " ")
}
