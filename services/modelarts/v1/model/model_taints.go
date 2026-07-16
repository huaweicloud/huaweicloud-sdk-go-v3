package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Taints 污点。
type Taints struct {

	// **参数解释**：键。 **取值范围**：不涉及。
	Key string `json:"key"`

	// **参数解释**：值。 **取值范围**：不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**：作用效果。 **取值范围**：不涉及。
	Effect string `json:"effect"`
}

func (o Taints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Taints struct{}"
	}

	return strings.Join([]string{"Taints", string(data)}, " ")
}
