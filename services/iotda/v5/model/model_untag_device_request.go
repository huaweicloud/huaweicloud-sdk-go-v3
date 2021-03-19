package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UntagDeviceRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *UnbindTagsDto `json:"body,omitempty"`
}

func (o UntagDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UntagDeviceRequest struct{}"
	}

	return strings.Join([]string{"UntagDeviceRequest", string(data)}, " ")
}
