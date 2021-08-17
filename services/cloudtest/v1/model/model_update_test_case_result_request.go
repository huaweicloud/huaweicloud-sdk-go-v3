package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTestCaseResultRequest struct {
	Body *UpdateTestCaseResultRequestBody `json:"body,omitempty"`
}

func (o UpdateTestCaseResultRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTestCaseResultRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseResultRequest", string(data)}, " ")
}
