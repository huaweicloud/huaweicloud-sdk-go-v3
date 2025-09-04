package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachVolumesRequest Request Object
type BatchDetachVolumesRequest struct {

	// 卷ID。
	VolumeId string `json:"volume_id"`

	Body *VolumeBatchDetachRequest `json:"body,omitempty"`
}

func (o BatchDetachVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachVolumesRequest struct{}"
	}

	return strings.Join([]string{"BatchDetachVolumesRequest", string(data)}, " ")
}
