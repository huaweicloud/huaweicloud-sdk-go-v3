package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopTaskRequest Request Object
type StopTaskRequest struct {

	// 迁移任务ID。
	TaskId string `json:"task_id"`
}

func (o StopTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTaskRequest struct{}"
	}

	return strings.Join([]string{"StopTaskRequest", string(data)}, " ")
}
