package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNodeResponse Response Object
type CreateNodeResponse struct {

	// **参数解释**： API类型 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： Node
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API类型 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： Node
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec,omitempty"`

	Status         *NodeStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodeResponse struct{}"
	}

	return strings.Join([]string{"CreateNodeResponse", string(data)}, " ")
}
