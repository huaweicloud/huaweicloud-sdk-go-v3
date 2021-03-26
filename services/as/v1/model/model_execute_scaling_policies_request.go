package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExecuteScalingPoliciesRequest struct {
	Body *ExecuteScalingPoliciesRequestBody `json:"body,omitempty"`
}

func (o ExecuteScalingPoliciesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPoliciesRequest", string(data)}, " ")
}
