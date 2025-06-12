package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResizeVolumeRequestBody This is a auto create Body Object
type BatchResizeVolumeRequestBody struct {

	// The list of disks to be expanded.
	Volumes []VolumesForBatchResizeVolume `json:"volumes"`

	BssParam *PrepaidParamForBatchResizeVolume `json:"bss_param,omitempty"`
}

func (o BatchResizeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"BatchResizeVolumeRequestBody", string(data)}, " ")
}
