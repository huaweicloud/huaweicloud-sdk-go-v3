package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRemuxTaskRequest struct {
	Body *CreateRemuxTaskReq `json:"body,omitempty"`
}

func (o CreateRemuxTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskRequest", string(data)}, " ")
}
