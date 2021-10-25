package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteProtectionGroupResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProtectionGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectionGroupResponse", string(data)}, " ")
}
