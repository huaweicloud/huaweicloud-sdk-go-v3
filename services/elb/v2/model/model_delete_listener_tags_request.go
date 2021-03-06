package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteListenerTagsRequest struct {
	// 监听器ID

	ListenerId string `json:"listener_id"`
	// 待删除标签的key值。

	Key string `json:"key"`
}

func (o DeleteListenerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerTagsRequest", string(data)}, " ")
}
