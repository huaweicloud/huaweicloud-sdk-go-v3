package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AttachProtectedInstanceReplicationResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachProtectedInstanceReplicationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AttachProtectedInstanceReplicationResponse struct{}"
	}

	return strings.Join([]string{"AttachProtectedInstanceReplicationResponse", string(data)}, " ")
}
