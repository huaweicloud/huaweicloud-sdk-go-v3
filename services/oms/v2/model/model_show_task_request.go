package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTaskRequest struct {
	// 任务ID

	TaskId int64 `json:"task_id"`
}

func (o ShowTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRequest", string(data)}, " ")
}
