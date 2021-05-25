package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteJobRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
}

func (o DeleteJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobRequest", string(data)}, " ")
}
