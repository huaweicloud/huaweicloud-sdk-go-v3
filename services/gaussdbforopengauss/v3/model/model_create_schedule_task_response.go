package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleTaskResponse Response Object
type CreateScheduleTaskResponse struct {

	// 任务ID。
	InstanceTasks  *[]InstanceTaskDetail `json:"instance_tasks,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CreateScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateScheduleTaskResponse", string(data)}, " ")
}
