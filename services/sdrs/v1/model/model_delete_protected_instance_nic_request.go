package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteProtectedInstanceNicRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *ProtectedInstanceDeleteNicRequestBody `json:"body,omitempty"`
}

func (o DeleteProtectedInstanceNicRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceNicRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceNicRequest", string(data)}, " ")
}
