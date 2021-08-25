package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListInstancesByResourceTagsRequest struct {
	Body *ListInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInstancesByResourceTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesByResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesByResourceTagsRequest", string(data)}, " ")
}
