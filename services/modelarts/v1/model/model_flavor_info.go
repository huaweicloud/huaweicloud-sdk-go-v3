package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorInfo 资源规格详细信息。
type FlavorInfo struct {

	// 可以选择的最大节点数量（max_num，为1代表不支持分布式）。
	MaxNum *int32 `json:"max_num,omitempty"`

	Cpu *Cpu `json:"cpu,omitempty"`

	Gpu *Gpu `json:"gpu,omitempty"`

	Npu *Npu `json:"npu,omitempty"`

	Memory *Memory `json:"memory,omitempty"`

	Disk *Disk `json:"disk,omitempty"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
