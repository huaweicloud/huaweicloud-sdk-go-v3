package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTestCaseRequest struct {
	Body *CreateTestCaseRequestBody `json:"body,omitempty"`
}

func (o CreateTestCaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTestCaseRequest struct{}"
	}

	return strings.Join([]string{"CreateTestCaseRequest", string(data)}, " ")
}
