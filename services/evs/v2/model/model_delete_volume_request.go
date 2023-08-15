package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVolumeRequest Request Object
type DeleteVolumeRequest struct {

	// 云硬盘ID。
	VolumeId string `json:"volume_id"`
}

func (o DeleteVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeRequest struct{}"
	}

	return strings.Join([]string{"DeleteVolumeRequest", string(data)}, " ")
}
