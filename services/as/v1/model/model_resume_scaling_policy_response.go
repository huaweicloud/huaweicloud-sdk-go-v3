package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResumeScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumeScalingPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResumeScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ResumeScalingPolicyResponse", string(data)}, " ")
}
