package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProtectedInstanceTagsRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`
}

func (o ListProtectedInstanceTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProtectedInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedInstanceTagsRequest", string(data)}, " ")
}
