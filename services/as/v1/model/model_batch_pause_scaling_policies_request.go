package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchPauseScalingPoliciesRequest struct {
	Body *BatchPauseScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchPauseScalingPoliciesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchPauseScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchPauseScalingPoliciesRequest", string(data)}, " ")
}
