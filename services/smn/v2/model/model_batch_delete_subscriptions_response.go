package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubscriptionsResponse Response Object
type BatchDeleteSubscriptionsResponse struct {

	// 取消订阅结果列表。
	SubscriptionsResult *[]BatchDeleteSubscriptionsResponseItemInfo `json:"subscriptions_result,omitempty"`
	HttpStatusCode      int                                         `json:"-"`
}

func (o BatchDeleteSubscriptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsResponse", string(data)}, " ")
}
