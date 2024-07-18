package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AsicAcceleratorInfo struct {

	// ASIC设备名称。
	Name *string `json:"name,omitempty"`

	// ASIC设备数量。
	Count *int32 `json:"count,omitempty"`

	// ASIC设备的内存，单位为MB。
	MemoryMb *int32 `json:"memory_mb,omitempty"`
}

func (o AsicAcceleratorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsicAcceleratorInfo struct{}"
	}

	return strings.Join([]string{"AsicAcceleratorInfo", string(data)}, " ")
}
