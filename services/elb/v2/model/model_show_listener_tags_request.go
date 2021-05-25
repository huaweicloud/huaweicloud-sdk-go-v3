package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowListenerTagsRequest struct {
	// 监听器ID。

	ListenerId string `json:"listener_id"`
}

func (o ShowListenerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowListenerTagsRequest", string(data)}, " ")
}
