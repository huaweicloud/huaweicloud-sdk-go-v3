package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateProtectedInstancesRequest struct {
	Body *BatchCreateProtectedInstancesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateProtectedInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedInstancesRequest", string(data)}, " ")
}
