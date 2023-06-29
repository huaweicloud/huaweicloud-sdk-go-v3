package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionSourceRequest Request Object
type UpdateSubscriptionSourceRequest struct {

	// 事件订阅ID
	SubscriptionId string `json:"subscription_id"`

	// 事件订阅源ID
	SourceId string `json:"source_id"`

	Body *SubscriptionSource `json:"body,omitempty"`
}

func (o UpdateSubscriptionSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionSourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionSourceRequest", string(data)}, " ")
}
