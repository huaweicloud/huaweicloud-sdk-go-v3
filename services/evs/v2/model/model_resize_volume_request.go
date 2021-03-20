package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResizeVolumeRequest struct {
	VolumeId string `json:"volume_id"`

	Body *ResizeVolumeRequestBody `json:"body,omitempty"`
}

func (o ResizeVolumeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequest struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequest", string(data)}, " ")
}
