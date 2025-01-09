package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduleTaskRequest Request Object
type UpdateScheduleTaskRequest struct {

	// 定时任务唯一标识。
	TaskId string `json:"task_id"`

	Body *UpdateScheduleTaskReq `json:"body,omitempty"`
}

func (o UpdateScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduleTaskRequest", string(data)}, " ")
}
