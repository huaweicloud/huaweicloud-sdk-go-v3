package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorView struct {
	FlavorId *FlavorId `json:"flavor_id,omitempty"`
	// 存储大小。

	StorageSize *string `json:"storage_size,omitempty"`
	// CPU限制。

	NumCpu *string `json:"num_cpu,omitempty"`
	// CPU初始。

	NumCpuInit *string `json:"num_cpu_init,omitempty"`
	// 内存限制。

	MemorySize *string `json:"memory_size,omitempty"`
	// 内存初始。

	MemorySizeInit *string `json:"memory_size_init,omitempty"`
	// 展示标签。

	Label *string `json:"label,omitempty"`
}

func (o FlavorView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorView struct{}"
	}

	return strings.Join([]string{"FlavorView", string(data)}, " ")
}
