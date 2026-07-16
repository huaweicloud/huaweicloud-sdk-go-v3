package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeMetadata **参数解释**：节点metadata信息。
type NodeMetadata struct {

	// **参数解释**：节点名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	Labels *NodeLabels `json:"labels,omitempty"`

	Annotations *NodeVoAnnotations `json:"annotations,omitempty"`
}

func (o NodeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeMetadata struct{}"
	}

	return strings.Join([]string{"NodeMetadata", string(data)}, " ")
}
