package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestDingdingEndpointInfo struct {

	// 终端地址。必须是一个钉钉群机器人的地址。
	Endpoint string `json:"endpoint"`

	// dingding协议订阅用户的关键字。dingding协议订阅用户必须指定keyword和sign_secret二者之一。当用户在钉钉机器人端添加关键字校验的安全策略时，这里的关键字必须是钉钉端所填写的关键字之一。
	Keyword *string `json:"keyword,omitempty"`

	// dingding协议订阅用户的加签密钥字段。dingding协议订阅用户必须指定keyword和sign_secret二者之一。当用户在钉钉机器人端添加关键字校验的安全策略时，这里的关键字必须是钉钉端所填写的关键字之一。
	SignSecret *string `json:"sign_secret,omitempty"`
}

func (o CreateSubscriptionUserRequestDingdingEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestDingdingEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestDingdingEndpointInfo", string(data)}, " ")
}
