package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResumeScalingGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumeScalingGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResumeScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"ResumeScalingGroupResponse", string(data)}, " ")
}
