package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeVolumeRequest Request Object
type ResizeVolumeRequest struct {

	// 云硬盘ID。
	VolumeId string `json:"volume_id"`

	Body *ResizeVolumeRequestBody `json:"body,omitempty"`
}

func (o ResizeVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequest struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequest", string(data)}, " ")
}
