package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExecuteScalingPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteScalingPoliciesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPoliciesResponse", string(data)}, " ")
}
