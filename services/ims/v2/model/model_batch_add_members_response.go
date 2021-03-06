package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchAddMembersResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddMembersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchAddMembersResponse", string(data)}, " ")
}
