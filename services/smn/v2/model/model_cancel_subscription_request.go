package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSubscriptionRequest Request Object
type CancelSubscriptionRequest struct {

	// 订阅者的唯一的资源标识，可通过[查询订阅者列表](ListSubscriptions.xml)获取该标识。
	SubscriptionUrn string `json:"subscription_urn"`
}

func (o CancelSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CancelSubscriptionRequest", string(data)}, " ")
}
