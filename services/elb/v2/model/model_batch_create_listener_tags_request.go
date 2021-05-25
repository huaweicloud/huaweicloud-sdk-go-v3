package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateListenerTagsRequest struct {
	// 监听器ID。

	ListenerId string `json:"listener_id"`

	Body *BatchCreateListenerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateListenerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsRequest", string(data)}, " ")
}
