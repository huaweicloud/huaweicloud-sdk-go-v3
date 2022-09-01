package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateVolumeTagsRequest struct {

	// 云硬盘ID。
	VolumeId string `json:"volume_id" xml:"volume_id"`

	Body *BatchCreateVolumeTagsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchCreateVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateVolumeTagsRequest", string(data)}, " ")
}
