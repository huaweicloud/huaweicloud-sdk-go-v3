package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMediaProcessTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o DeleteMediaProcessTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMediaProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMediaProcessTaskRequest", string(data)}, " ")
}
