package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceConfig **参数解释**：作业启动资源配置。  **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type ResourceConfig struct {

	// **参数解释**：CPU数量（单位：毫核）  **约束限制**：不涉及。 **取值范围**：300~100000，包含边界值。 **默认取值**：不涉及。
	Cpu int32 `json:"cpu"`

	// **参数解释**：内存配置（单位：兆字节）  **约束限制**：不涉及。 **取值范围**：500~204800，包含边界值。 **默认取值**：不涉及。
	Memory int32 `json:"memory"`

	// **参数解释**：NPU配置，卡数。  **约束限制**：不涉及。 **取值范围**：1~8，包含边界值。 **默认取值**：不涉及。
	Npu *int32 `json:"npu,omitempty"`

	// **参数解释**：GPU配置，卡数。  **约束限制**：不涉及。 **取值范围**：1~8，包含边界值。 **默认取值**：不涉及。
	Gpu *int32 `json:"gpu,omitempty"`
}

func (o ResourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceConfig struct{}"
	}

	return strings.Join([]string{"ResourceConfig", string(data)}, " ")
}
