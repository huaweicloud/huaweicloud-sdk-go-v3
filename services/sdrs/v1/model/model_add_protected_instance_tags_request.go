package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddProtectedInstanceTagsRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *ProtectedInstanceAddTagsRequestBody `json:"body,omitempty"`
}

func (o AddProtectedInstanceTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddProtectedInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"AddProtectedInstanceTagsRequest", string(data)}, " ")
}
