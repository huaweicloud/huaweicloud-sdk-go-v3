package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetUserPasswordRequest struct {
	// DDM实例ID。

	InstanceId string `json:"instance_id"`
	// 需要修改的DDM帐号名称。

	Username string `json:"username"`

	Body *ResetUserPasswordReq `json:"body,omitempty"`
}

func (o ResetUserPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetUserPasswordRequest", string(data)}, " ")
}
