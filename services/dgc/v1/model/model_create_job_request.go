package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateJobRequest struct {
	Body *JobInfo `json:"body,omitempty"`
}

func (o CreateJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateJobRequest struct{}"
	}

	return strings.Join([]string{"CreateJobRequest", string(data)}, " ")
}
