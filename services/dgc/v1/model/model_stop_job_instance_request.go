package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StopJobInstanceRequest struct {
	JobName    string `json:"job_name"`
	InstanceId string `json:"instance_id"`
}

func (o StopJobInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopJobInstanceRequest", string(data)}, " ")
}
