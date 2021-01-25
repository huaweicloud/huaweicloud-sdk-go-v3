package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ModifyResourceRequest struct {
	ResourceId string        `json:"resource_id"`
	Body       *ResourceInfo `json:"body,omitempty"`
}

func (o ModifyResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyResourceRequest struct{}"
	}

	return strings.Join([]string{"ModifyResourceRequest", string(data)}, " ")
}
