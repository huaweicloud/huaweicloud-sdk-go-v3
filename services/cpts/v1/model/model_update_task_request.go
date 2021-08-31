package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTaskRequest struct {
	// 任务id

	TaskId int32 `json:"task_id"`

	Body *UpdateTaskRequestBody `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
