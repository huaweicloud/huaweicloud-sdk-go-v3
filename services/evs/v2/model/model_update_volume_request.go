package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateVolumeRequest struct {
	// 云硬盘ID。

	VolumeId string `json:"volume_id"`

	Body *UpdateVolumeRequestBody `json:"body,omitempty"`
}

func (o UpdateVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeRequest struct{}"
	}

	return strings.Join([]string{"UpdateVolumeRequest", string(data)}, " ")
}
