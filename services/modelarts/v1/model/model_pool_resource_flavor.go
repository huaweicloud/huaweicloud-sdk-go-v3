package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolResourceFlavor 资源池规格数据模型。
type PoolResourceFlavor struct {

	// **参数解释**：资源规格，比如：modelarts.vm.gpu.tnt004。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Flavor string `json:"flavor"`

	// **参数解释**：资源规格的保障资源量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Count int32 `json:"count"`

	// **参数解释**：资源规格的弹性资源量。物理池中该值和count必须一致。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	MaxCount *int32 `json:"maxCount,omitempty"`

	ExtendParams *PoolResourceFlavorExtendParams `json:"extendParams,omitempty"`

	Os *Os `json:"os,omitempty"`
}

func (o PoolResourceFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolResourceFlavor struct{}"
	}

	return strings.Join([]string{"PoolResourceFlavor", string(data)}, " ")
}
