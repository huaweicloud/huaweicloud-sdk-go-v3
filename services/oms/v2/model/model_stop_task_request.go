package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StopTaskRequest struct {
	// 迁移任务ID。

	TaskId int64 `json:"task_id"`
}

func (o StopTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopTaskRequest struct{}"
	}

	return strings.Join([]string{"StopTaskRequest", string(data)}, " ")
}
