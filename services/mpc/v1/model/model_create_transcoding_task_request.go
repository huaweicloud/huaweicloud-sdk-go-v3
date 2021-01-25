package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTranscodingTaskRequest struct {
	Body *CreateTranscodingReq `json:"body,omitempty"`
}

func (o CreateTranscodingTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTranscodingTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTranscodingTaskRequest", string(data)}, " ")
}
