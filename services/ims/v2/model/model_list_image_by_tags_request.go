package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListImageByTagsRequest struct {
	Body *ListImageByTagsRequestBody `json:"body,omitempty"`
}

func (o ListImageByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListImageByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImageByTagsRequest", string(data)}, " ")
}
