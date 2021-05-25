package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchPauseScalingPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchPauseScalingPoliciesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchPauseScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchPauseScalingPoliciesResponse", string(data)}, " ")
}
