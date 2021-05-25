package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
}

func (o ShowJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
