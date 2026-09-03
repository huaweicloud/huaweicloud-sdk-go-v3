package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestFeishuEndpointInfo struct {

	// 终端地址。必须是一个飞书群机器人的地址。
	Endpoint string `json:"endpoint"`

	// feishu协议订阅用户的关键字。feishu协议订阅用户必须指定keyword和sign_secret二者之一。当用户在飞书机器人端添加关键字校验的安全策略时，这里的关键字必须是飞书端所填写的关键字之一。
	Keyword *string `json:"keyword,omitempty"`

	// feishu协议订阅用户的加签密钥字段。feishu协议订阅用户必须指定keyword和sign_secret二者之一。当用户在飞书机器人端添加关键字校验的安全策略时，这里的关键字必须是飞书端所填写的关键字之一。
	SignSecret *string `json:"sign_secret,omitempty"`
}

func (o CreateSubscriptionUserRequestFeishuEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestFeishuEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestFeishuEndpointInfo", string(data)}, " ")
}
