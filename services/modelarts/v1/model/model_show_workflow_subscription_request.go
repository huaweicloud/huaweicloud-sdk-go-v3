package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowSubscriptionRequest Request Object
type ShowWorkflowSubscriptionRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流的消息订阅ID。
	SubscriptionId string `json:"subscription_id"`
}

func (o ShowWorkflowSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowSubscriptionRequest", string(data)}, " ")
}
