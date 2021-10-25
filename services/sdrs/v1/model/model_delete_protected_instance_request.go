package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteProtectedInstanceRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *DeleteProtectedInstanceRequestBody `json:"body,omitempty"`
}

func (o DeleteProtectedInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceRequest", string(data)}, " ")
}
