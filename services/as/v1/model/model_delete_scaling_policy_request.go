package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteScalingPolicyRequest struct {
	// 伸缩策略ID。

	ScalingPolicyId string `json:"scaling_policy_id"`
}

func (o DeleteScalingPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingPolicyRequest", string(data)}, " ")
}
