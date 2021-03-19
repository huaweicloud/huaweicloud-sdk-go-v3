package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddDeviceRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddDevice `json:"body,omitempty"`
}

func (o AddDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddDeviceRequest struct{}"
	}

	return strings.Join([]string{"AddDeviceRequest", string(data)}, " ")
}
