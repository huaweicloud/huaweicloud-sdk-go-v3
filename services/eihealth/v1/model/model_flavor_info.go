package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorInfo notebook规格信息
type FlavorInfo struct {

	// **参数解释**： cpu架构类型。 **约束限制**： 不涉及 **取值范围**： * X86: X86架构 * ARM：ARM架构 **默认取值**： X86
	CpuType *string `json:"cpu_type,omitempty"`

	// notebook占用的cpu,0.1核为100m,单位为\"C\"
	Cpu float32 `json:"cpu"`

	// notebook占用的gpu，0.1为使用单卡10%，1为占满单个显卡，1+为使用多个显卡
	Gpu float32 `json:"gpu"`

	// **参数解释**： gpu架构类型。 **约束限制**： 不涉及 **取值范围**： * GPU：支持GPU * [Snt9：支持Snt9](tag:hws) **默认取值**： 不涉及
	GpuType *string `json:"gpu_type,omitempty"`

	// notebook占用的内存,单位为\"G\"
	Memory float32 `json:"memory"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
