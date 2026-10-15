package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RayHeadResourceSpec struct {

	// **参数解释**：CPU核数量，单位为毫核。 **约束限制**：不涉及。 **取值范围**：最小值为1000。 **默认取值**：不涉及。
	Cpu int32 `json:"cpu"`

	// **参数解释**：单个head内存大小，单位MB。 **约束限制**：不涉及。 **取值范围**：最小值为1024。 **默认取值**：不涉及。
	Memory int32 `json:"memory"`

	// **参数解释**：节点的GPU数量。 **约束限制**：不支持小于1的值或带小数部分的值（如0.5, 1.5）。 **取值范围**：需配置为大于等于1.0，且必须为整数值（如1.0, 2.0, 3.0, ...）。 **默认取值**：不涉及。
	Gpu *float64 `json:"gpu,omitempty"`

	// **参数解释**：节点的NPU数量。 **约束限制**：不支持小于1的值或带小数部分的值（如0.5, 1.5）。 **取值范围**：需配置为大于等于1.0，且必须为整数值（如1.0, 2.0, 3.0, ...）。 **默认取值**：不涉及。
	Npu *float64 `json:"npu,omitempty"`
}

func (o RayHeadResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayHeadResourceSpec struct{}"
	}

	return strings.Join([]string{"RayHeadResourceSpec", string(data)}, " ")
}
