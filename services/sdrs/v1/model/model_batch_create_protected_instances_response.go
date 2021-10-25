package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateProtectedInstancesResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateProtectedInstancesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedInstancesResponse", string(data)}, " ")
}
