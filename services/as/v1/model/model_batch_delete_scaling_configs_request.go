package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteScalingConfigsRequest struct {
	Body *BatchDeleteScalingConfigOption `json:"body,omitempty"`
}

func (o BatchDeleteScalingConfigsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingConfigsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingConfigsRequest", string(data)}, " ")
}
