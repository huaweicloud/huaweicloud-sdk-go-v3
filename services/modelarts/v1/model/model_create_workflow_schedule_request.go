package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowScheduleRequest Request Object
type CreateWorkflowScheduleRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	Body *WorkflowSchedule `json:"body,omitempty"`
}

func (o CreateWorkflowScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowScheduleRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowScheduleRequest", string(data)}, " ")
}
