package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateScalingV2PolicyRequest struct {
	// 伸缩策略ID。

	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *UpdateScalingV2PolicyOption `json:"body,omitempty"`
}

func (o UpdateScalingV2PolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateScalingV2PolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateScalingV2PolicyRequest", string(data)}, " ")
}
