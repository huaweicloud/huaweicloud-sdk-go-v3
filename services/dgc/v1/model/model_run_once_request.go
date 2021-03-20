package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunOnceRequest struct {
	JobName string `json:"job_name"`

	Body *StartJobReq `json:"body,omitempty"`
}

func (o RunOnceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunOnceRequest struct{}"
	}

	return strings.Join([]string{"RunOnceRequest", string(data)}, " ")
}
