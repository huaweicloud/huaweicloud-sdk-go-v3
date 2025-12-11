package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostProperties struct {

	// 核数。
	CpuCores *int32 `json:"cpu_cores,omitempty"`

	// 内存。
	MemoryMb *int32 `json:"memory_mb,omitempty"`

	// 主频。
	CpuSpeed *string `json:"cpu_speed,omitempty"`
}

func (o HostProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostProperties struct{}"
	}

	return strings.Join([]string{"HostProperties", string(data)}, " ")
}
