package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GpuInfo struct {

	// GPU设备名称。
	Name *string `json:"name,omitempty"`

	// GPU设备数量。
	Count *int32 `json:"count,omitempty"`

	// GPU设备的内存，单位为MB。
	MemoryMb *int32 `json:"memory_mb,omitempty"`
}

func (o GpuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GpuInfo struct{}"
	}

	return strings.Join([]string{"GpuInfo", string(data)}, " ")
}
