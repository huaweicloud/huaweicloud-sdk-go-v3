package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ImportJobResponse struct {
	// 任务id
	TaskId         *string `json:"taskId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportJobResponse struct{}"
	}

	return strings.Join([]string{"ImportJobResponse", string(data)}, " ")
}
