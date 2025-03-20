package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScheduledTaskRequest Request Object
type ShowScheduledTaskRequest struct {

	// ID of ScheduledTask
	TaskId string `json:"task_id"`
}

func (o ShowScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowScheduledTaskRequest", string(data)}, " ")
}
