package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeTopicRequest Request Object
type SubscribeTopicRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。当订阅类型为短信或邮件时，与token、endpoint参数组合使用，三个参数中设置任意两个或全部设置均可确认订阅成功；当订阅类型为其他时，该参数必填，需与token同时设置。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 订阅终端地址。仅当订阅类型为短信或邮件时可设置，与token、topic_urn参数组合使用，三个参数中设置任意两个或全部设置均可确认订阅成功。
	Endpoint *string `json:"endpoint,omitempty"`

	// 确认订阅Token信息。（订阅链接中携带的token信息）
	Token string `json:"token"`
}

func (o SubscribeTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeTopicRequest struct{}"
	}

	return strings.Join([]string{"SubscribeTopicRequest", string(data)}, " ")
}
