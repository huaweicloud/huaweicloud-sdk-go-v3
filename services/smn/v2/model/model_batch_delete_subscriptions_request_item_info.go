package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteSubscriptionsRequestItemInfo struct {

	// 订阅者的唯一的资源标识。
	SubscriptionUrn string `json:"subscription_urn"`
}

func (o BatchDeleteSubscriptionsRequestItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsRequestItemInfo struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsRequestItemInfo", string(data)}, " ")
}
