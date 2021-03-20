package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateResourceRequest struct {
	ResourceId string `json:"resource_id"`

	Body *ResourceInfo `json:"body,omitempty"`
}

func (o UpdateResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateResourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceRequest", string(data)}, " ")
}
