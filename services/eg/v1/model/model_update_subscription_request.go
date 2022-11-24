package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSubscriptionRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`

	Body *SubscriptionUpdateReq `json:"body,omitempty"`
}

func (o UpdateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionRequest", string(data)}, " ")
}
