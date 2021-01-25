package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateTagRequest struct {
	ResourceType string         `json:"resource_type"`
	ResourceId   string         `json:"resource_id"`
	Body         *BatchHandTags `json:"body,omitempty"`
}

func (o BatchCreateTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateTagRequest", string(data)}, " ")
}
