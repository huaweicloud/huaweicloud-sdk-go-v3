package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEventRequest struct {
	// 事件ID

	Eventid string `json:"eventid"`
}

func (o ShowEventRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEventRequest struct{}"
	}

	return strings.Join([]string{"ShowEventRequest", string(data)}, " ")
}
