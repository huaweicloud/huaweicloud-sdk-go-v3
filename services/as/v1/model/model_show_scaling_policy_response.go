package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowScalingPolicyResponse struct {
	ScalingPolicy  *ScalingV1PolicyDetail `json:"scaling_policy,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowScalingPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowScalingPolicyResponse", string(data)}, " ")
}
