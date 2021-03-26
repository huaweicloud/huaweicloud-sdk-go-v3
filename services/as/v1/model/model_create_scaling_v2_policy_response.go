package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateScalingV2PolicyResponse struct {
	// 伸缩策略ID。

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateScalingV2PolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateScalingV2PolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingV2PolicyResponse", string(data)}, " ")
}
