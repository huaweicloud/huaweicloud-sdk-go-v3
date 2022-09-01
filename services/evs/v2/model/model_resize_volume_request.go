package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResizeVolumeRequest struct {

	// 云硬盘ID。
	VolumeId string `json:"volume_id" xml:"volume_id"`

	Body *ResizeVolumeRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ResizeVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequest struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequest", string(data)}, " ")
}
