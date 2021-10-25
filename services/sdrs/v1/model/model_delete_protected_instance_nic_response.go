package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteProtectedInstanceNicResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProtectedInstanceNicResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceNicResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceNicResponse", string(data)}, " ")
}
