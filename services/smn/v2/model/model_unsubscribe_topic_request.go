package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeTopicRequest Request Object
type UnsubscribeTopicRequest struct {

	// 订阅者的唯一的资源标识，可通过[查询订阅者列表](ListSubscriptions.xml)获取该标识。
	SubscriptionUrn string `json:"subscription_urn"`
}

func (o UnsubscribeTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeTopicRequest struct{}"
	}

	return strings.Join([]string{"UnsubscribeTopicRequest", string(data)}, " ")
}
