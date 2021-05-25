package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateVolumeRequest struct {
	// 云硬盘ID。

	VolumeId string `json:"volume_id"`

	Body *UpdateVolumeRequestBody `json:"body,omitempty"`
}

func (o UpdateVolumeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVolumeRequest struct{}"
	}

	return strings.Join([]string{"UpdateVolumeRequest", string(data)}, " ")
}
