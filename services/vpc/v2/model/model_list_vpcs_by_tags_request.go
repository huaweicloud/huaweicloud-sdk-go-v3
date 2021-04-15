package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVpcsByTagsRequest struct {
	Body *ListVpcsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListVpcsByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVpcsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcsByTagsRequest", string(data)}, " ")
}
