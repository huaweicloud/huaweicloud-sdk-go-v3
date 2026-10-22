package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceCapacityItem **参数解释**：容量项与资源ID映射。 **取值范围**：不涉及。
type ResourceCapacityItem struct {

	// **参数解释**：资源类型。 **约束限制**：不涉及。 **取值范围**：   - CPU：中央处理器。   - GPU：图形处理器。   - NPU：神经网络处理器。 **默认取值**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：资源容量数量。CPU资源单位为核，GPU/NPU资源单位为个。 **约束限制**：不涉及。 **取值范围**：0~48000。 **默认取值**：不涉及。
	Capacity *int32 `json:"capacity,omitempty"`

	// **参数解释**：规格编码。 **取值范围**：长度为1~64的英文字母、数字、中划线和点号的组合。
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**：资源ID。 **取值范围**：长度为1~64的英文字母、数字、中划线和点号的组合。
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o ResourceCapacityItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCapacityItem struct{}"
	}

	return strings.Join([]string{"ResourceCapacityItem", string(data)}, " ")
}
