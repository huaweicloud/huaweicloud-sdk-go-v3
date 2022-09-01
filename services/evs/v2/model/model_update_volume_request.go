package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateVolumeRequest struct {

	// 云硬盘ID。
	VolumeId string `json:"volume_id" xml:"volume_id"`

	Body *UpdateVolumeRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeRequest struct{}"
	}

	return strings.Join([]string{"UpdateVolumeRequest", string(data)}, " ")
}
