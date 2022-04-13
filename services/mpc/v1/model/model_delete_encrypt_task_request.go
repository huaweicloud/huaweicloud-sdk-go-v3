package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEncryptTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o DeleteEncryptTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteEncryptTaskRequest", string(data)}, " ")
}
