package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePool 节点池的详细信息。
type NodePool struct {

	// **参数解释**： API版本。 **取值范围**： 可选值如下： - v2
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：节点池类型。 **取值范围**： 可选值如下： - NodePool：节点池
	Kind string `json:"kind"`

	Metadata *NodePoolMetadata `json:"metadata"`

	Spec *NodePoolSpec `json:"spec"`

	Status *NodePoolStatus `json:"status,omitempty"`
}

func (o NodePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePool struct{}"
	}

	return strings.Join([]string{"NodePool", string(data)}, " ")
}
