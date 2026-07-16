package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodePoolNodesResponse Response Object
type ListNodePoolNodesResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v2：当前资源版本为v2。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - NodeList：节点列表。
	Kind *string `json:"kind,omitempty"`

	Metadata *NodeListMetadata `json:"metadata,omitempty"`

	// **参数解释**：节点资源列表。
	Items          *[]Node `json:"items,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNodePoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodePoolNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNodePoolNodesResponse", string(data)}, " ")
}
