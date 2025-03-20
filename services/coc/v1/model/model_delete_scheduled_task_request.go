package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduledTaskRequest Request Object
type DeleteScheduledTaskRequest struct {

	// ID of ScheduledTask
	TaskId string `json:"task_id"`
}

func (o DeleteScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteScheduledTaskRequest", string(data)}, " ")
}
