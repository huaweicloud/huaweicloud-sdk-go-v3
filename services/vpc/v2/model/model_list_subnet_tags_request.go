package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSubnetTagsRequest struct {
}

func (o ListSubnetTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetTagsRequest", string(data)}, " ")
}
