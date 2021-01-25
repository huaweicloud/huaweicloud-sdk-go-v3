package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobStatusRequest struct {
	JobName string `json:"job_name"`
}

func (o ShowJobStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowJobStatusRequest", string(data)}, " ")
}
