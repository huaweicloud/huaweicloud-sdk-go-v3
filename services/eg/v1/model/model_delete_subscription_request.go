package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSubscriptionRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`
}

func (o DeleteSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionRequest", string(data)}, " ")
}
