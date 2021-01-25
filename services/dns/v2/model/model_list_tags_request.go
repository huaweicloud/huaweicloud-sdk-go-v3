package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTagsRequest struct {
	ResourceType string `json:"resource_type"`
}

func (o ListTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}
