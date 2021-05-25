package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExecuteScalingPolicyRequest struct {
	// 伸缩策略ID。

	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *ExecuteScalingPolicyOption `json:"body,omitempty"`
}

func (o ExecuteScalingPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyRequest", string(data)}, " ")
}
