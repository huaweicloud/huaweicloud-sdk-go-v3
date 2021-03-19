package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMediaProcessTaskResponse struct {
	// 任务Id

	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMediaProcessTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessTaskResponse", string(data)}, " ")
}
