package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSubscriptionTargetRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id" xml:"subscription_id"`

	// 事件订阅目标ID
	TargetId string `json:"target_id" xml:"target_id"`

	Body *SubscriptionTarget `json:"body,omitempty" xml:"body"`
}

func (o UpdateSubscriptionTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionTargetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionTargetRequest", string(data)}, " ")
}
