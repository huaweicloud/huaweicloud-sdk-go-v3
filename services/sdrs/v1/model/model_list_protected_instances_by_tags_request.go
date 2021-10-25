package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProtectedInstancesByTagsRequest struct {
	Body *ListProtectedInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListProtectedInstancesByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesByTagsRequest", string(data)}, " ")
}
