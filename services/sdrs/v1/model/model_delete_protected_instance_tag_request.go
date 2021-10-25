package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteProtectedInstanceTagRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`
	// 标签key。

	Key string `json:"key"`
}

func (o DeleteProtectedInstanceTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceTagRequest", string(data)}, " ")
}
