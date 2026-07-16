package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeSelector struct {

	// **参数解释**：必填项。节点选择器项的列表。这些项是“或”的关系。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms"`
}

func (o NodeSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSelector struct{}"
	}

	return strings.Join([]string{"NodeSelector", string(data)}, " ")
}
