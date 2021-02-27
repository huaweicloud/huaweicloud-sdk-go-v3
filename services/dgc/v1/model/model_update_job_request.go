package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateJobRequest struct {
	JobName string   `json:"job_name"`
	Body    *JobInfo `json:"body,omitempty"`
}

func (o UpdateJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobRequest", string(data)}, " ")
}
