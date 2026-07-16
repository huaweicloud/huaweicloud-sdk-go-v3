package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowSubscriptionRequest Request Object
type UpdateWorkflowSubscriptionRequest struct {

	// 消息订阅ID。
	SubscriptionId string `json:"subscription_id"`

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	Body *Subscription `json:"body,omitempty"`
}

func (o UpdateWorkflowSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowSubscriptionRequest", string(data)}, " ")
}
