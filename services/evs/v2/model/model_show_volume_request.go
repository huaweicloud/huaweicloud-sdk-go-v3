package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVolumeRequest struct {
	// 云硬盘ID。

	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeRequest", string(data)}, " ")
}
