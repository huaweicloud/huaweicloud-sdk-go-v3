package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEncryptTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteEncryptTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteEncryptTaskRequest", string(data)}, " ")
}
