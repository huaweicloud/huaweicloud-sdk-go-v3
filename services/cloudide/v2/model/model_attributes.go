package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Attributes struct {
	// cpu限制

	CpuLimit *string `json:"cpu_limit,omitempty"`
	// 内存限制

	MemoryLimitBytes *string `json:"memory_limit_bytes,omitempty"`
	// pvc规格

	PvcQuantity *string `json:"pvc_quantity,omitempty"`
}

func (o Attributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attributes struct{}"
	}

	return strings.Join([]string{"Attributes", string(data)}, " ")
}
