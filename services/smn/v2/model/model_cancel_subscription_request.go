package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelSubscriptionRequest struct {
	// 订阅者的唯一的资源标识，可通过[查询订阅者列表](https://support.huaweicloud.com/api-smn/ListSubscriptions.html)获取该标识。

	SubscriptionUrn string `json:"subscription_urn"`
}

func (o CancelSubscriptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CancelSubscriptionRequest", string(data)}, " ")
}
