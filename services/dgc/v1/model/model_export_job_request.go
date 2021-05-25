package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExportJobRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
}

func (o ExportJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportJobRequest struct{}"
	}

	return strings.Join([]string{"ExportJobRequest", string(data)}, " ")
}
