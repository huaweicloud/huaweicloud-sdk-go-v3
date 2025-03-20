package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledTaskRequest Request Object
type UpdateScheduledTaskRequest struct {

	// ID of ScheduledTask
	TaskId string `json:"task_id"`

	Body *ScheduledTaskRequestBody `json:"body,omitempty"`
}

func (o UpdateScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTaskRequest", string(data)}, " ")
}
