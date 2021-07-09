package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateUsersRequest struct {
	// DDM实例ID。

	InstanceId string `json:"instance_id"`

	Body *CreateUsersReq `json:"body,omitempty"`
}

func (o CreateUsersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUsersRequest struct{}"
	}

	return strings.Join([]string{"CreateUsersRequest", string(data)}, " ")
}
