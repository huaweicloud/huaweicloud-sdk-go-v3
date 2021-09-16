package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchRebootInstanceRequest struct {
	Body *BatchRebootInstanceRequestBody `json:"body,omitempty"`
}

func (o BatchRebootInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchRebootInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootInstanceRequest", string(data)}, " ")
}
