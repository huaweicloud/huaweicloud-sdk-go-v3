package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowServerBlockDeviceResponse struct {
	VolumeAttachment *ServerBlockDevice `json:"volumeAttachment,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ShowServerBlockDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowServerBlockDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowServerBlockDeviceResponse", string(data)}, " ")
}
