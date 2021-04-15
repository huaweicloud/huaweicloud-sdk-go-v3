package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSubnetsByTagsRequest struct {
	Body *ListSubnetsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListSubnetsByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubnetsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetsByTagsRequest", string(data)}, " ")
}
