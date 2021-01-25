package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateEncryptTaskResponse struct {
	// 加密任务Id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEncryptTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEncryptTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateEncryptTaskResponse", string(data)}, " ")
}
