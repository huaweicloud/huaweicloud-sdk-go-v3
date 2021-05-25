package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateScalingGroupRequest struct {
	Body *CreateScalingGroupOption `json:"body,omitempty"`
}

func (o CreateScalingGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupRequest", string(data)}, " ")
}
