package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCommandRequest struct {
	DeviceId string `json:"device_id"`

	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *DeviceCommandRequest `json:"body,omitempty"`
}

func (o CreateCommandRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCommandRequest struct{}"
	}

	return strings.Join([]string{"CreateCommandRequest", string(data)}, " ")
}
