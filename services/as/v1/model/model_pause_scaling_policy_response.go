package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type PauseScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseScalingPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PauseScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"PauseScalingPolicyResponse", string(data)}, " ")
}
