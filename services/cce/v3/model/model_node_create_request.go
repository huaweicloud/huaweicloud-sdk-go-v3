package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeCreateRequest
type NodeCreateRequest struct {

	// **参数解释**： API类型 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： Node
	Kind string `json:"kind"`

	// **参数解释**： API版本 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： v3
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
