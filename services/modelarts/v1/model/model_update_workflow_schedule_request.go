package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowScheduleRequest Request Object
type UpdateWorkflowScheduleRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流调度信息ID。
	ScheduleId string `json:"schedule_id"`

	Body *WorkflowScheduleUpdate `json:"body,omitempty"`
}

func (o UpdateWorkflowScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowScheduleRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowScheduleRequest", string(data)}, " ")
}
