package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowScalingV2PolicyRequest struct {
	// 伸缩组ID。

	ScalingPolicyId string `json:"scaling_policy_id"`
}

func (o ShowScalingV2PolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowScalingV2PolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingV2PolicyRequest", string(data)}, " ")
}
