package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type TagDeviceRequest struct {
	InstanceId *string      `json:"Instance-Id,omitempty"`
	Body       *BindTagsDto `json:"body,omitempty"`
}

func (o TagDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TagDeviceRequest struct{}"
	}

	return strings.Join([]string{"TagDeviceRequest", string(data)}, " ")
}
