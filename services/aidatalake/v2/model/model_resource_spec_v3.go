package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSpecV3 **参数解释**：资源配置。
type ResourceSpecV3 struct {

	// **参数解释**：资源规格。 **取值范围**：不涉及。
	SpecCode string `json:"spec_code"`

	// **参数解释**：CPU核数量，单位为毫核。 **约束限制**：不涉及。 **取值范围**：0~1000000。 **默认取值**：0。
	Cpu int32 `json:"cpu"`

	// **参数解释**：内存大小，单位MB。 **约束限制**：不涉及。 **取值范围**：0~10000000。 **默认取值**：0。
	Memory int32 `json:"memory"`

	// **参数解释**：昇腾卡数量。 **约束限制**：暂时只支持整数。 **取值范围**：0~100。 **默认取值**：0。
	Npu *float64 `json:"npu,omitempty"`

	// **参数解释**：GPU卡数量。 **约束限制**：暂时只支持整数。 **取值范围**：0~100。 **默认取值**：0。
	Gpu *float64 `json:"gpu,omitempty"`
}

func (o ResourceSpecV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpecV3 struct{}"
	}

	return strings.Join([]string{"ResourceSpecV3", string(data)}, " ")
}
