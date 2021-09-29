package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListServerTagsRequest struct {
}

func (o ListServerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListServerTagsRequest", string(data)}, " ")
}
