package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceUserRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *CreateInstanceUserReq `json:"body,omitempty"`
}

func (o CreateInstanceUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceUserRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceUserRequest", string(data)}, " ")
}
