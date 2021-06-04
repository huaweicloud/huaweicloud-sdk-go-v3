package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateSecurityGroupResponse struct {
	// 工作流ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSecurityGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupResponse", string(data)}, " ")
}
