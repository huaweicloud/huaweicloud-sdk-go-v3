package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumesForBatchResizeVolume struct {

	// The disk ID.
	Id string `json:"id"`

	// The new disk size, in GiB. This parameter value must be greater than the original disk size and less than the maximum size allowed for a disk. The maximum disk size: - Data disk: 32,768 GiB - System disk: 1,024 GiB
	NewSize int32 `json:"new_size"`
}

func (o VolumesForBatchResizeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumesForBatchResizeVolume struct{}"
	}

	return strings.Join([]string{"VolumesForBatchResizeVolume", string(data)}, " ")
}
