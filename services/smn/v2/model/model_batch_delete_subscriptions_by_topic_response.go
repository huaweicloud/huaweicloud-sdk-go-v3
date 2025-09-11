package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubscriptionsByTopicResponse Response Object
type BatchDeleteSubscriptionsByTopicResponse struct {

	// 取消订阅结果列表。
	SubscriptionsResult *[]BatchDeleteSubscriptionsResponseItemInfo `json:"subscriptions_result,omitempty"`
	HttpStatusCode      int                                         `json:"-"`
}

func (o BatchDeleteSubscriptionsByTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsByTopicResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsByTopicResponse", string(data)}, " ")
}
