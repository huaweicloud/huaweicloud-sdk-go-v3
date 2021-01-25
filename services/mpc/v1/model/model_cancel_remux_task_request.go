package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelRemuxTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o CancelRemuxTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelRemuxTaskRequest", string(data)}, " ")
}
