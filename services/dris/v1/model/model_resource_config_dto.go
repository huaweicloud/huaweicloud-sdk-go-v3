package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceConfigDto struct {

	// **参数说明**：cpu个数。
	Cpu *float32 `json:"cpu,omitempty"`

	// **参数说明**：内存大小。
	Memory *float32 `json:"memory,omitempty"`

	// **参数说明**：gpu个数。
	Gpu *float32 `json:"gpu,omitempty"`

	// **参数说明**：npu个数。
	Npu *float32 `json:"npu,omitempty"`
}

func (o ResourceConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceConfigDto struct{}"
	}

	return strings.Join([]string{"ResourceConfigDto", string(data)}, " ")
}
