package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteTagsRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *BatchDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequest", string(data)}, " ")
}
