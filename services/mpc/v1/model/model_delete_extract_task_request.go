package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteExtractTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o DeleteExtractTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteExtractTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteExtractTaskRequest", string(data)}, " ")
}
