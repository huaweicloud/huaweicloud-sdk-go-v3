package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StartJobRequest struct {
	JobName string `json:"job_name"`

	Body *StartJobReq `json:"body,omitempty"`
}

func (o StartJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartJobRequest struct{}"
	}

	return strings.Join([]string{"StartJobRequest", string(data)}, " ")
}
