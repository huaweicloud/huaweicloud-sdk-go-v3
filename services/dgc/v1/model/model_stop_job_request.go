package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StopJobRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
}

func (o StopJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopJobRequest struct{}"
	}

	return strings.Join([]string{"StopJobRequest", string(data)}, " ")
}
