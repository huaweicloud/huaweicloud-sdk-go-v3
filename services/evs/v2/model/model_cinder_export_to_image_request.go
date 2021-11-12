package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CinderExportToImageRequest struct {
	// 云硬盘ID。

	VolumeId string `json:"volume_id"`

	Body *CinderExportToImageRequestBody `json:"body,omitempty"`
}

func (o CinderExportToImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderExportToImageRequest struct{}"
	}

	return strings.Join([]string{"CinderExportToImageRequest", string(data)}, " ")
}
