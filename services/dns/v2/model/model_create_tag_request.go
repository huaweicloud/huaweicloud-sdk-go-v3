package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTagRequest struct {
	ResourceType string `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	Body *CreateTagReq `json:"body,omitempty"`
}

func (o CreateTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTagRequest", string(data)}, " ")
}
