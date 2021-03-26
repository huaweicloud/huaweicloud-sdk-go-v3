package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateScalingV2PolicyRequest struct {
	Body *CreateScalingPolicyV2RequestBody `json:"body,omitempty"`
}

func (o CreateScalingV2PolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateScalingV2PolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingV2PolicyRequest", string(data)}, " ")
}
