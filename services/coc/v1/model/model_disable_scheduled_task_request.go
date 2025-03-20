package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableScheduledTaskRequest Request Object
type DisableScheduledTaskRequest struct {

	// ID of ScheduledTask
	TaskId string `json:"task_id"`
}

func (o DisableScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"DisableScheduledTaskRequest", string(data)}, " ")
}
