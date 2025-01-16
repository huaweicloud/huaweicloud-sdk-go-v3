package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionExtension struct {

	// 这是应用ID字段。当protocol值为welink时，该字段为必填字段，从welink方获取。
	ClientId *string `json:"client_id,omitempty"`

	// 该字段为应用secret字段。当protocol值为welink时，该字段为必填字段，从welink方获取。
	ClientSecret *string `json:"client_secret,omitempty"`

	// 该字段为关键字字段。当protocol值为feishu时，这里的keyword字段和sign_secret字段二者必选其一。当用户在飞书或钉钉自定义机器人端添加关键字校验的安全策略时，这里的关键字必须是飞书或钉钉自定义机器人中所填写的关键字之一。
	Keyword *string `json:"keyword,omitempty"`

	// 这是加签密钥字段。当protocol协议为feishu时，这个字段和keyword字段二者必选且只能选其一。密钥配置必须与客户在飞书或钉钉自定义机器人的密钥配置完全一致。例如，如果在飞书端配置了密钥并且没有配置关键字，则在此处填入从飞书获取的密钥字段，如果在飞书端没有配置密钥并且配置了关键字，则不填写该字段。
	SignSecret *string `json:"sign_secret,omitempty"`

	// 该字段为http header字段，用户可以在字段限制范围内自定义http header，header字段内容以KV对形式存在。当使用主题发送时，已确认的订阅发送消息会携带用户自定义的http header。 header应满足如下要求： key值限定为：包含英文字母([A-Za-z])、数字([0-9])、中划线(-)hyphens，中划线不得作为结尾，不得连续出现。 K/V不得超过10个 key需要以\"x-\"开头，不能以\"x-smn\"开头，正确示例：x-abc-cba, x-abc 所有K/V长度总和不得超过1024个字符 key不区分大小写 key值不可重复 value值限定为ASCII码，不支持中文或其他Unicode，支持空格
	Header map[string]string `json:"header,omitempty"`

	// 个人钉钉appKey字段，字符长度限制64个，仅支持字母、数字、中划线(-)、下划线(_)。当订阅协议为dingTalkBot时，该字段必选。
	AppKey *string `json:"app_key,omitempty"`

	// 个人钉钉appSecret字段，字符长度限制128个，仅支持字母、数字、中划线(-)、下划线(_)。当订阅协议为dingTalkBot时，该字段必选。
	AppSecret *string `json:"app_secret,omitempty"`

	// 个人钉钉robotCode字段，名称：机器人编码，字符长度限制64个，仅支持字母、数字、中划线(-)、下划线(_)，一般与appKey一致。当订阅协议为dingTalkBot时，该字段必选。
	RobotCode *string `json:"robot_code,omitempty"`
}

func (o SubscriptionExtension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionExtension struct{}"
	}

	return strings.Join([]string{"SubscriptionExtension", string(data)}, " ")
}
