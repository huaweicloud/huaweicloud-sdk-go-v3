package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTranscodingTaskRequest struct {
	// 创建转码任务成功时返回的任务ID

	TaskId int32 `json:"task_id"`
}

func (o DeleteTranscodingTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskRequest", string(data)}, " ")
}
