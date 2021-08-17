package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunTestCaseRequest struct {
	Body *RunTestCaseRequestBody `json:"body,omitempty"`
}

func (o RunTestCaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunTestCaseRequest struct{}"
	}

	return strings.Join([]string{"RunTestCaseRequest", string(data)}, " ")
}
