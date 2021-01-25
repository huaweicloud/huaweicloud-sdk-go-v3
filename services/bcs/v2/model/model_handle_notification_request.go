package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type HandleNotificationRequest struct {
	Body *HandleNotificationRequestBody `json:"body,omitempty"`
}

func (o HandleNotificationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HandleNotificationRequest struct{}"
	}

	return strings.Join([]string{"HandleNotificationRequest", string(data)}, " ")
}
