package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteListenerRequest struct {
	// 监听器ID。

	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteListenerRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerRequest", string(data)}, " ")
}
