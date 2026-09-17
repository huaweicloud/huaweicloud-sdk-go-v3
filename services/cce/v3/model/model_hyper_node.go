package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HyperNode **参数解释**： 超节点 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type HyperNode struct {

	// **参数解释**： API版本，固定值v3。 **约束限制**： 该值不可修改 **取值范围**： - v3  **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： API类型，固定值HyperNode。 **约束限制**： 该值不可修改 **取值范围**： - HyperNode  **默认取值**： HyperNode
	Kind *string `json:"kind,omitempty"`

	Metadata *HyperNodeMetadata `json:"metadata,omitempty"`

	Spec *HyperNodeSpec `json:"spec,omitempty"`

	Status *HyperNodeStatus `json:"status,omitempty"`
}

func (o HyperNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNode struct{}"
	}

	return strings.Join([]string{"HyperNode", string(data)}, " ")
}
