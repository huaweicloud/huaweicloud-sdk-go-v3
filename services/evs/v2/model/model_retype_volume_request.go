package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetypeVolumeRequest Request Object
type RetypeVolumeRequest struct {

	// 磁盘ID。
	VolumeId string `json:"volume_id"`

	Body *RetypeVolumeRequestBody `json:"body,omitempty"`
}

func (o RetypeVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetypeVolumeRequest struct{}"
	}

	return strings.Join([]string{"RetypeVolumeRequest", string(data)}, " ")
}
