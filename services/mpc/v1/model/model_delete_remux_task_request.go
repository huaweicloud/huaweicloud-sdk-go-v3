package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRemuxTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteRemuxTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteRemuxTaskRequest", string(data)}, " ")
}
