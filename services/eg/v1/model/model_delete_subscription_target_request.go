package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionTargetRequest Request Object
type DeleteSubscriptionTargetRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`

	// 事件订阅目标ID
	TargetId string `json:"target_id"`
}

func (o DeleteSubscriptionTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionTargetRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionTargetRequest", string(data)}, " ")
}
