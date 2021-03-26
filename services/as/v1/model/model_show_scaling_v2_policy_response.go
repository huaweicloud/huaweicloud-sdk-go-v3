package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowScalingV2PolicyResponse struct {
	ScalingPolicy  *ScalingV2PolicyDetail `json:"scaling_policy,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowScalingV2PolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowScalingV2PolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowScalingV2PolicyResponse", string(data)}, " ")
}
