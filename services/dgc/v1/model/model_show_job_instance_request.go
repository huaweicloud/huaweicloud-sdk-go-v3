package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobInstanceRequest struct {
	JobName    string `json:"job_name"`
	InstanceId string `json:"instance_id"`
}

func (o ShowJobInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowJobInstanceRequest", string(data)}, " ")
}
