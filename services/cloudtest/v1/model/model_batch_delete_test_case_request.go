package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteTestCaseRequest struct {
	Body *BatchDeleteTestCaseRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteTestCaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCaseRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCaseRequest", string(data)}, " ")
}
