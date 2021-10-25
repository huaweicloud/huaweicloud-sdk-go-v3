package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AddProtectedInstanceNicResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddProtectedInstanceNicResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddProtectedInstanceNicResponse struct{}"
	}

	return strings.Join([]string{"AddProtectedInstanceNicResponse", string(data)}, " ")
}
