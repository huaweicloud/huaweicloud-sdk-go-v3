package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopTaskByIdRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o StopTaskByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTaskByIdRequest struct{}"
	}

	return strings.Join([]string{"StopTaskByIdRequest", string(data)}, " ")
}
