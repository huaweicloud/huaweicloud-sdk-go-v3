package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ModifyJobRequest struct {
	JobName string   `json:"job_name"`
	Body    *JobInfo `json:"body,omitempty"`
}

func (o ModifyJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyJobRequest struct{}"
	}

	return strings.Join([]string{"ModifyJobRequest", string(data)}, " ")
}
