package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateListenerTagsRequest struct {
	ListenerId string `json:"listener_id"`

	Body *CreateListenerTagsRequestBody `json:"body,omitempty"`
}

func (o CreateListenerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsRequest", string(data)}, " ")
}
