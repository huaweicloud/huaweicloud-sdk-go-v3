package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPropertiesRequest struct {
	DeviceId       string  `json:"device_id"`
	StageAuthToken *string `json:"Stage-Auth-Token,omitempty"`
	InstanceId     *string `json:"Instance-Id,omitempty"`
	ServiceId      string  `json:"service_id"`
}

func (o ListPropertiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPropertiesRequest struct{}"
	}

	return strings.Join([]string{"ListPropertiesRequest", string(data)}, " ")
}
