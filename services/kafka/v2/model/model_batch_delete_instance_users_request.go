package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteInstanceUsersRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *BatchDeleteInstanceUsersReq `json:"body,omitempty"`
}

func (o BatchDeleteInstanceUsersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceUsersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceUsersRequest", string(data)}, " ")
}
