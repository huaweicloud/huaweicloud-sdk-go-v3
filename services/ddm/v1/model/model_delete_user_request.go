package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteUserRequest struct {
	// DDM实例ID。

	InstanceId string `json:"instance_id"`
	// 要删除的DDM帐号名称。

	Username string `json:"username"`
}

func (o DeleteUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserRequest", string(data)}, " ")
}
