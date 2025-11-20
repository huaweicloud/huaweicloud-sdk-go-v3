package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodesResponse Response Object
type ListNodesResponse struct {

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： List
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 固定值,不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 节点对象列表，包含了当前集群下所有节点的详细信息。可通过items.metadata.name下的值来找到对应的节点。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Items          *[]Node `json:"items,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNodesResponse", string(data)}, " ")
}
