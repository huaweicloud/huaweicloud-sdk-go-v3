package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateScalingPolicyRequest struct {
	// 伸缩策略ID。

	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *UpdateScalingPolicyOption `json:"body,omitempty"`
}

func (o UpdateScalingPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateScalingPolicyRequest", string(data)}, " ")
}
