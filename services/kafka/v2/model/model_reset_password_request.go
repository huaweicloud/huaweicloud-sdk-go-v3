package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetPasswordRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ResetPasswordReq `json:"body,omitempty"`
}

func (o ResetPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequest", string(data)}, " ")
}
