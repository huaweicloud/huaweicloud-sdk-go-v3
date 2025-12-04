package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeTopicRequest Request Object
type SubscribeTopicRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 订阅终端地址。
	Endpoint *string `json:"endpoint,omitempty"`

	// 订阅主题Token信息。
	Token string `json:"token"`
}

func (o SubscribeTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeTopicRequest struct{}"
	}

	return strings.Join([]string{"SubscribeTopicRequest", string(data)}, " ")
}
