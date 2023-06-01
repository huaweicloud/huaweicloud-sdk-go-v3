package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSubscriptionRequest struct {

	// Topic的唯一的资源标识。可以通过[查看主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 订阅者的唯一的资源标识，可通过[查询订阅者列表](ListSubscriptions.xml)获取该标识。
	SubscriptionUrn string `json:"subscription_urn"`

	Body *UpdateSubscriptionRequestBody `json:"body,omitempty"`
}

func (o UpdateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionRequest", string(data)}, " ")
}
