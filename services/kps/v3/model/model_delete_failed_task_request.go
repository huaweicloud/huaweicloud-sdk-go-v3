package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteFailedTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o DeleteFailedTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFailedTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteFailedTaskRequest", string(data)}, " ")
}
