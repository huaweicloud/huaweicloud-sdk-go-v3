package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTaskRequest struct {
	Body *CreateTaskReq `json:"body,omitempty"`
}

func (o CreateTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskRequest", string(data)}, " ")
}
