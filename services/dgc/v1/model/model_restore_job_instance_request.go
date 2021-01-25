package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreJobInstanceRequest struct {
	JobName    string `json:"job_name"`
	InstanceId string `json:"instance_id"`
}

func (o RestoreJobInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreJobInstanceRequest", string(data)}, " ")
}
