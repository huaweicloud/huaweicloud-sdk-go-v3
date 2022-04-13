package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteFailedTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o DeleteFailedTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFailedTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteFailedTaskRequest", string(data)}, " ")
}
