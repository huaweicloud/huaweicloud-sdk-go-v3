package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchResumeScalingPoliciesRequest struct {
	Body *BatchResumeScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchResumeScalingPoliciesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchResumeScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchResumeScalingPoliciesRequest", string(data)}, " ")
}
