package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type PauseScalingGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseScalingGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PauseScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"PauseScalingGroupResponse", string(data)}, " ")
}
