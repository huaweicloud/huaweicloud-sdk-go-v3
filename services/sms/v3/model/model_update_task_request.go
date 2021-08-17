package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTaskRequest struct {
	// 迁移任务id

	TaskId string `json:"task_id"`

	Body *PutTaskReq `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
