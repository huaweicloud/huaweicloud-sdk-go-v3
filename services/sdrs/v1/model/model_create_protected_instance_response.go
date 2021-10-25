package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateProtectedInstanceResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProtectedInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProtectedInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateProtectedInstanceResponse", string(data)}, " ")
}
