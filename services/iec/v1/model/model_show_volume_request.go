package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVolumeRequest struct {

	// 磁盘ID。
	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeRequest", string(data)}, " ")
}
