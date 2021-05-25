package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowInstanceUsersRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceUsersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceUsersRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceUsersRequest", string(data)}, " ")
}
