package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartReverseProtectionGroupResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartReverseProtectionGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartReverseProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"StartReverseProtectionGroupResponse", string(data)}, " ")
}
