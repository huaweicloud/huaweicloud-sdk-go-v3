package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowNewTaskStatusRequest struct {
	// 任务ID（非负整数）的字符串

	TaskId *string `json:"task_id,omitempty"`
}

func (o ShowNewTaskStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowNewTaskStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowNewTaskStatusRequest", string(data)}, " ")
}
