package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScheduleTaskRequest Request Object
type ShowScheduleTaskRequest struct {

	// 定时任务唯一标识。
	TaskId string `json:"task_id"`
}

func (o ShowScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowScheduleTaskRequest", string(data)}, " ")
}
