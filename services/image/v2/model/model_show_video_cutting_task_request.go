package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVideoCuttingTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowVideoCuttingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoCuttingTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoCuttingTaskRequest", string(data)}, " ")
}
