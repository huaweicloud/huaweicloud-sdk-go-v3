package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Node
type Node struct {

	// **参数解释**： API类型，固定值“Node”。 **取值范围**： 只能为固定值“Node”。
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本，固定值“v3”。 **取值范围**： 只能为固定值“v3”。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec,omitempty"`

	Status *NodeStatus `json:"status,omitempty"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}
