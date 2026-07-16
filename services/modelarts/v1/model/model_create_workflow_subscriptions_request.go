package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowSubscriptionsRequest Request Object
type CreateWorkflowSubscriptionsRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	Body *Subscription `json:"body,omitempty"`
}

func (o CreateWorkflowSubscriptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowSubscriptionsRequest", string(data)}, " ")
}
