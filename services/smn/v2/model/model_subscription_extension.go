package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionExtension struct {

	// 该字段为welink订阅下的租户ID字段，由租户从welink方获取。当protocol值为welink时，该字段为必填字段。
	ClientId *string `json:"client_id,omitempty"`

	// 该字段为welink订阅下的租户获取的client secret字段，由租户从welink方获取。当protocol值为welink时，该字段为必填字段。
	ClientSecret *string `json:"client_secret,omitempty"`

	// 该字段为关键字字段。当protocol协议为feishu时，这里的keyword字段和sign_secret字段二者必选其一。当用户在飞书或钉钉机器人端添加关键字校验的安全策略时，这里的关键字必须是飞书或钉钉端所填写的关键字之一。
	Keyword *string `json:"keyword,omitempty"`

	// 这是加签密钥字段。当protocol协议为feishu或dingding时，这个字段和keyword字段二者必选且只能选其一，密钥配置必须与客户在飞书或钉钉客户端的密钥配置完全一致。例如，如果在飞书端配置了密钥并且没有配置关键字，则在此处填入从飞书获取的密钥字段，如果在飞书端没有配置密钥并且配置了关键字，则不填写该字段。
	SignSecret *string `json:"sign_secret,omitempty"`
}

func (o SubscriptionExtension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionExtension struct{}"
	}

	return strings.Join([]string{"SubscriptionExtension", string(data)}, " ")
}
