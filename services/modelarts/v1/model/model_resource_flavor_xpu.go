package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorXpu 算力卡信息，包含类型、卡数、单卡显存等
type ResourceFlavorXpu struct {

	// **参数解释**：卡类型。 **取值范围**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：芯片数量。reseverd for backwards compatibility **取值范围**：不涉及。
	Size *string `json:"size,omitempty"`

	// **参数解释**：单卡显存大小。 **取值范围**：不涉及。
	Memory *string `json:"memory,omitempty"`

	// **参数解释**：卡数量。 **取值范围**：不涉及。
	Card *string `json:"card,omitempty"`

	// **参数解释**：芯片数量。值同size字段一致。 **取值范围**：不涉及。
	Chip *string `json:"chip,omitempty"`
}

func (o ResourceFlavorXpu) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorXpu struct{}"
	}

	return strings.Join([]string{"ResourceFlavorXpu", string(data)}, " ")
}
