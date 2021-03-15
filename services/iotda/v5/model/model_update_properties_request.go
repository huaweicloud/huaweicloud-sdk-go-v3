package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePropertiesRequest struct {
	DeviceId       string                   `json:"device_id"`
	StageAuthToken *string                  `json:"Stage-Auth-Token,omitempty"`
	InstanceId     *string                  `json:"Instance-Id,omitempty"`
	Body           *DevicePropertiesRequest `json:"body,omitempty"`
}

func (o UpdatePropertiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePropertiesRequest struct{}"
	}

	return strings.Join([]string{"UpdatePropertiesRequest", string(data)}, " ")
}
