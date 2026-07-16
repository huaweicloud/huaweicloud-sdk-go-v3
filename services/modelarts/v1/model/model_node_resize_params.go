package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeResizeParams 节点规格变更参数
type NodeResizeParams struct {

	// **参数解释**：节点池。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodePool *string `json:"nodePool,omitempty"`

	// **参数解释**：规格。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：步长。 **约束限制**：不涉及。
	CreatingStep *interface{} `json:"creatingStep,omitempty"`
}

func (o NodeResizeParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResizeParams struct{}"
	}

	return strings.Join([]string{"NodeResizeParams", string(data)}, " ")
}
