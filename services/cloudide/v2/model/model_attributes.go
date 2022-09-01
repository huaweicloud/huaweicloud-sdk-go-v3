package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Attributes struct {

	// cpu限制
	CpuLimit *string `json:"cpu_limit,omitempty" xml:"cpu_limit"`

	// 内存限制
	MemoryLimitBytes *string `json:"memory_limit_bytes,omitempty" xml:"memory_limit_bytes"`

	// pvc规格
	PvcQuantity *string `json:"pvc_quantity,omitempty" xml:"pvc_quantity"`
}

func (o Attributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attributes struct{}"
	}

	return strings.Join([]string{"Attributes", string(data)}, " ")
}
