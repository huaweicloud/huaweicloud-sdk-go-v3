package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableScheduledTaskRequest Request Object
type EnableScheduledTaskRequest struct {

	// ID of ScheduledTask
	TaskId string `json:"task_id"`

	Body *EnableScheduledTaskRequestBody `json:"body,omitempty"`
}

func (o EnableScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"EnableScheduledTaskRequest", string(data)}, " ")
}
