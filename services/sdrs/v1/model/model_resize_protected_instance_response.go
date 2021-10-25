package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResizeProtectedInstanceResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeProtectedInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResizeProtectedInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeProtectedInstanceResponse", string(data)}, " ")
}
