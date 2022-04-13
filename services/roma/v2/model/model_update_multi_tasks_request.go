package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMultiTasksRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 任务ID

	TaskId string `json:"task_id"`

	Body *MultiTaskUpdateBody `json:"body,omitempty"`
}

func (o UpdateMultiTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMultiTasksRequest struct{}"
	}

	return strings.Join([]string{"UpdateMultiTasksRequest", string(data)}, " ")
}
