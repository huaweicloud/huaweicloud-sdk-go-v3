package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResizeVolumesRequest Request Object
type BatchResizeVolumesRequest struct {
	Body *BatchResizeVolumeRequestBody `json:"body,omitempty"`
}

func (o BatchResizeVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeVolumesRequest struct{}"
	}

	return strings.Join([]string{"BatchResizeVolumesRequest", string(data)}, " ")
}
