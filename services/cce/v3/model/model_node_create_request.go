package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeCreateRequest
type NodeCreateRequest struct {

	// **参数解释**： API类型，固定值“Node”。 **约束限制**： 不涉及 **取值范围**： 只能为固定值“Node”。 **默认取值**： 不涉及
	Kind string `json:"kind"`

	// **参数解释**： API版本，固定值“v3”。 **约束限制**： 不涉及 **取值范围**： 只能为固定值“v3”。 **默认取值**： 不涉及
	ApiVersion string `json:"apiVersion"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec"`
}

func (o NodeCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeCreateRequest struct{}"
	}

	return strings.Join([]string{"NodeCreateRequest", string(data)}, " ")
}
