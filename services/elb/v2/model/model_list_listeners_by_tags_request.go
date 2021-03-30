package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListListenersByTagsRequest struct {
	Body *ListListenersByTagsRequestBody `json:"body,omitempty"`
}

func (o ListListenersByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListListenersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsRequest", string(data)}, " ")
}
