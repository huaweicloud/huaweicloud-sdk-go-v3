package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowResourceTagRequest struct {
	ResourceType string `json:"resource_type"`

	ResourceId string `json:"resource_id"`
}

func (o ShowResourceTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowResourceTagRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagRequest", string(data)}, " ")
}
