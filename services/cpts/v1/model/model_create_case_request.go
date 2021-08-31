package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCaseRequest struct {
	Body *CreateCaseRequestBody `json:"body,omitempty"`
}

func (o CreateCaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCaseRequest struct{}"
	}

	return strings.Join([]string{"CreateCaseRequest", string(data)}, " ")
}
