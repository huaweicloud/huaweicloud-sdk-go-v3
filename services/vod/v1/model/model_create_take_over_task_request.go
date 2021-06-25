package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTakeOverTaskRequest struct {
	Body *CreateTakeOverTaskReq `json:"body,omitempty"`
}

func (o CreateTakeOverTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTakeOverTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTakeOverTaskRequest", string(data)}, " ")
}
