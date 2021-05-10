package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAllTagsRequest struct {
}

func (o ListAllTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAllTagsRequest struct{}"
	}

	return strings.Join([]string{"ListAllTagsRequest", string(data)}, " ")
}
