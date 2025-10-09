package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeResponse Response Object
type ShowNodeResponse struct {

	// **参数解释**： API类型，固定值“Node”。 **取值范围**： 只能为固定值“Node”。
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本，固定值“v3”。 **取值范围**： 只能为固定值“v3”。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec,omitempty"`

	Status         *NodeStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeResponse", string(data)}, " ")
}
