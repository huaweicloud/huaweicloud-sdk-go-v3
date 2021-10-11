package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchUpdateConfigsRequest struct {
	Body *BatchUpdateConfigs `json:"body,omitempty"`
}

func (o BatchUpdateConfigsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUpdateConfigsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateConfigsRequest", string(data)}, " ")
}
