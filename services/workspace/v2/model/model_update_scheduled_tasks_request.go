package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledTasksRequest Request Object
type UpdateScheduledTasksRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`

	Body *UpdateScheduledTasksReq `json:"body,omitempty"`
}

func (o UpdateScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTasksRequest", string(data)}, " ")
}
