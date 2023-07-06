package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopTaskByIdRequest Request Object
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
