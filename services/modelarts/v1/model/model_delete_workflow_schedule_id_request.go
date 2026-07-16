package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowScheduleIdRequest Request Object
type DeleteWorkflowScheduleIdRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流调度信息ID。
	ScheduleId string `json:"schedule_id"`
}

func (o DeleteWorkflowScheduleIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowScheduleIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowScheduleIdRequest", string(data)}, " ")
}
