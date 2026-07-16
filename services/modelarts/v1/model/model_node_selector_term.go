package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSelectorTerm **参数解释**：空节点选择器或空节点选择器项不匹配任何对象。 **约束限制**：要求是按“与”（AND）逻辑进行组合。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeSelectorTerm struct {

	// **参数解释**：按节点标签列出的节点选择器要求。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	MatchExpressions *[]NodeSelectorRequirement `json:"matchExpressions,omitempty"`

	// **参数解释**：按节点字段列出的节点选择器要求。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	MatchFields *[]NodeSelectorRequirement `json:"matchFields,omitempty"`
}

func (o NodeSelectorTerm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSelectorTerm struct{}"
	}

	return strings.Join([]string{"NodeSelectorTerm", string(data)}, " ")
}
