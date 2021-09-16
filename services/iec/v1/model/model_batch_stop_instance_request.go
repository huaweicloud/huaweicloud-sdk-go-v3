package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchStopInstanceRequest struct {
	Body *BatchStopInstanceRequestBody `json:"body,omitempty"`
}

func (o BatchStopInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchStopInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchStopInstanceRequest", string(data)}, " ")
}
