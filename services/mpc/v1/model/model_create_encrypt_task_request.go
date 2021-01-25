package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEncryptTaskRequest struct {
	Body *CreateEncryptReq `json:"body,omitempty"`
}

func (o CreateEncryptTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateEncryptTaskRequest", string(data)}, " ")
}
