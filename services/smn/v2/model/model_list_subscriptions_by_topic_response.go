package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubscriptionsByTopicResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 订阅者个数。
	SubscriptionCount *int32 `json:"subscription_count,omitempty" xml:"subscription_count"`

	// Subscription结构体。
	Subscriptions  *[]ListSubscriptionsItem `json:"subscriptions,omitempty" xml:"subscriptions"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListSubscriptionsByTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsByTopicResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsByTopicResponse", string(data)}, " ")
}
