package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyVolumeQoSRequest Request Object
type ModifyVolumeQoSRequest struct {

	// 云硬盘ID。
	VolumeId string `json:"volume_id"`

	Body *ModifyVolumeQoSRequestBody `json:"body,omitempty"`
}

func (o ModifyVolumeQoSRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVolumeQoSRequest struct{}"
	}

	return strings.Join([]string{"ModifyVolumeQoSRequest", string(data)}, " ")
}
