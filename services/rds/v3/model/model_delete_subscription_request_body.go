package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionRequestBody 删除订阅。
type DeleteSubscriptionRequestBody struct {

	// 删除的订阅id列表。每次删除的订阅必须属于同一实例。
	SubscriptionIds []string `json:"subscription_ids"`

	// 清理订阅配置。默认为false。  true：清理。 false：不清理。
	CleanSubscription *bool `json:"clean_subscription,omitempty"`
}

func (o DeleteSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionRequestBody", string(data)}, " ")
}
