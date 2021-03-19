package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTagRequest struct {
	ResourceType string `json:"resource_type"`

	Body *ListTagReq `json:"body,omitempty"`
}

func (o ListTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTagRequest struct{}"
	}

	return strings.Join([]string{"ListTagRequest", string(data)}, " ")
}
