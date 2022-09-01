package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceConfigDto struct {

	// cpu个数
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu"`

	// 内存大小
	Memory *float32 `json:"memory,omitempty" xml:"memory"`

	// cpu个数
	Gpu *float32 `json:"gpu,omitempty" xml:"gpu"`

	// cpu个数
	Npu *float32 `json:"npu,omitempty" xml:"npu"`
}

func (o ResourceConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceConfigDto struct{}"
	}

	return strings.Join([]string{"ResourceConfigDto", string(data)}, " ")
}
