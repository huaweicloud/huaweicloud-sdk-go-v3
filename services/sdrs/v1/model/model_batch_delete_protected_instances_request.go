package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteProtectedInstancesRequest struct {
	Body *BatchDeleteProtectedInstancesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteProtectedInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedInstancesRequest", string(data)}, " ")
}
