package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleTaskRequest Request Object
type DeleteScheduleTaskRequest struct {

	// 定时任务唯一标识。
	TaskId string `json:"task_id"`
}

func (o DeleteScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteScheduleTaskRequest", string(data)}, " ")
}
