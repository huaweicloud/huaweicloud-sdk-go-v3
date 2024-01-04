package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSubscriptionsFilterPolicesRequestBodyPolices struct {

	// 订阅者的唯一的资源标识，可通过[查询订阅者列表](ListSubscriptions.xml)获取该标识。
	SubscriptionUrn string `json:"subscription_urn"`

	// 订阅者的过滤策略。策略name不能重复
	FilterPolices []SubscriptionsFilterPolicy `json:"filter_polices"`
}

func (o BatchSubscriptionsFilterPolicesRequestBodyPolices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSubscriptionsFilterPolicesRequestBodyPolices struct{}"
	}

	return strings.Join([]string{"BatchSubscriptionsFilterPolicesRequestBodyPolices", string(data)}, " ")
}
