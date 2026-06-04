package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindPublicGatewayResponse Response Object
type UnbindPublicGatewayResponse struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：** 实例名称。 **取值范围：** 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释：** 节点ID。 **取值范围：** 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释：** 节点名称。 **取值范围：** 不涉及。
	NodeName       *string `json:"node_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnbindPublicGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindPublicGatewayResponse struct{}"
	}

	return strings.Join([]string{"UnbindPublicGatewayResponse", string(data)}, " ")
}
