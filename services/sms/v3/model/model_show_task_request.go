package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTaskRequest struct {
	// 迁移任务id

	TaskId string `json:"task_id"`
}

func (o ShowTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRequest", string(data)}, " ")
}
