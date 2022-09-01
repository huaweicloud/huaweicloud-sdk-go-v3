package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSubscriptionTargetRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id" xml:"subscription_id"`

	// 事件订阅目标ID
	TargetId string `json:"target_id" xml:"target_id"`
}

func (o DeleteSubscriptionTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionTargetRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionTargetRequest", string(data)}, " ")
}