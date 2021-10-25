package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateReplicationResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReplicationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateReplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateReplicationResponse", string(data)}, " ")
}
