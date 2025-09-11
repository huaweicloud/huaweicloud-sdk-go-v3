package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteSubscriptionsRequestBody struct {

	// 订阅者信息列表。
	SubscriptionUrns []BatchDeleteSubscriptionsRequestItemInfo `json:"subscription_urns"`
}

func (o BatchDeleteSubscriptionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsRequestBody", string(data)}, " ")
}
