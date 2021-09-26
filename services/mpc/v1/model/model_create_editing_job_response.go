package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateEditingJobResponse struct {
	// 接受任务后产生的任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEditingJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEditingJobResponse struct{}"
	}

	return strings.Join([]string{"CreateEditingJobResponse", string(data)}, " ")
}
