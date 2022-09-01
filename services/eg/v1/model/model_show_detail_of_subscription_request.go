package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailOfSubscriptionRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id" xml:"subscription_id"`
}

func (o ShowDetailOfSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfSubscriptionRequest", string(data)}, " ")
}
