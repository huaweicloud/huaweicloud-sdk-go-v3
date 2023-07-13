package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Res struct {

	// cpu数量，字符串所对应的数值在0.01到1000之间
	Cpu *string `json:"cpu,omitempty"`

	// gpu数量，值在0到1000
	Gpu *string `json:"gpu,omitempty"`

	// 内存数量，如果是资源限制，其值范围在4到1024000之间，否则在0.01到1024000之间
	Memory *string `json:"memory,omitempty"`

	// npu数量，字符串所对应的数值在0.到1000之间
	Npu *string `json:"npu,omitempty"`
}

func (o Res) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Res struct{}"
	}

	return strings.Join([]string{"Res", string(data)}, " ")
}
