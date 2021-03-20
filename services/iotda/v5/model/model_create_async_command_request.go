package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAsyncCommandRequest struct {
	DeviceId string `json:"device_id"`

	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AsyncDeviceCommandRequest `json:"body,omitempty"`
}

func (o CreateAsyncCommandRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAsyncCommandRequest struct{}"
	}

	return strings.Join([]string{"CreateAsyncCommandRequest", string(data)}, " ")
}
