package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNodePoolResponse Response Object
type CreateNodePoolResponse struct {

	// **参数解释**：API版本。 **取值范围**：可选值如下： - v2
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：节点池类型。 **取值范围**：可选值如下： -  NodePool：节点池
	Kind *string `json:"kind,omitempty"`

	Metadata *CreateNodePoolMetaVo `json:"metadata,omitempty"`

	Spec           *NodePoolSpec `json:"spec,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolResponse struct{}"
	}

	return strings.Join([]string{"CreateNodePoolResponse", string(data)}, " ")
}
