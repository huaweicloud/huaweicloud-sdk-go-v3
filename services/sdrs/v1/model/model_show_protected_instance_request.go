package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProtectedInstanceRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`
}

func (o ShowProtectedInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProtectedInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowProtectedInstanceRequest", string(data)}, " ")
}
