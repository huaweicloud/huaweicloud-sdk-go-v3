package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowListenerRequest struct {
	// 监听器ID。

	ListenerId string `json:"listener_id"`
}

func (o ShowListenerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowListenerRequest struct{}"
	}

	return strings.Join([]string{"ShowListenerRequest", string(data)}, " ")
}
