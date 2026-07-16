package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Node 节点的数据模型。
type Node struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v2：当前资源版本为v2。
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - Node：节点。
	Kind string `json:"kind"`

	Metadata *NodeMetadata `json:"metadata"`

	Spec *NodeSpec `json:"spec"`

	Status *NodeStatus `json:"status"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}
