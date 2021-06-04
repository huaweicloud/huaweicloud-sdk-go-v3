package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceNameRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceNameRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}
