package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateUserRequest struct {
	// DDM实例ID。

	InstanceId string `json:"instance_id"`
	// 需要修改的DDM帐号名称。

	Username string `json:"username"`

	Body *UpdateUserReq `json:"body,omitempty"`
}

func (o UpdateUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserRequest", string(data)}, " ")
}
