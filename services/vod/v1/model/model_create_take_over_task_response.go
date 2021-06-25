package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTakeOverTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTakeOverTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTakeOverTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTakeOverTaskResponse", string(data)}, " ")
}
