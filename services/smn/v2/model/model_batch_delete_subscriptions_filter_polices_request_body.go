package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteSubscriptionsFilterPolicesRequestBody struct {

	// 订阅者的唯一的资源标识，可通过[查询订阅者列表](ListSubscriptions.xml)获取该标识。
	SubscriptionUrns []string `json:"subscription_urns"`
}

func (o BatchDeleteSubscriptionsFilterPolicesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsFilterPolicesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsFilterPolicesRequestBody", string(data)}, " ")
}
