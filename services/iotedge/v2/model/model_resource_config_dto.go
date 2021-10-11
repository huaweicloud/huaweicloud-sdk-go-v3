package model

import (
	"encoding/json"

	"strings"
)

type ResourceConfigDto struct {
	// cpu个数

	Cpu *float32 `json:"cpu,omitempty"`
	// 内存大小

	Memory *float32 `json:"memory,omitempty"`
	// cpu个数

	Gpu *float32 `json:"gpu,omitempty"`
	// cpu个数

	Npu *float32 `json:"npu,omitempty"`
}

func (o ResourceConfigDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceConfigDto struct{}"
	}

	return strings.Join([]string{"ResourceConfigDto", string(data)}, " ")
}
