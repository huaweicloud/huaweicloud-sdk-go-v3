package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolNodeResponse Response Object
type ShowPoolNodeResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v2：当前资源版本为v2。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - Node：节点。
	Kind *string `json:"kind,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec,omitempty"`

	Status         *NodeStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPoolNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolNodeResponse", string(data)}, " ")
}
