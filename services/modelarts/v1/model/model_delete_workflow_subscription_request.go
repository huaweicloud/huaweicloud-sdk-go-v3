package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowSubscriptionRequest Request Object
type DeleteWorkflowSubscriptionRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流的消息订阅ID。
	SubscriptionId string `json:"subscription_id"`
}

func (o DeleteWorkflowSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowSubscriptionRequest", string(data)}, " ")
}
