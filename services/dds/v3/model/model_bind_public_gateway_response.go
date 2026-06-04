package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindPublicGatewayResponse Response Object
type BindPublicGatewayResponse struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：** 实例名称。 **取值范围：** 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释：** 节点ID。 **取值范围：** 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释：** 节点名称。 **取值范围：** 不涉及。
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释：** 公网NAT网关实例的ID。 **取值范围：** 不涉及。
	NatGatewayId *string `json:"nat_gateway_id,omitempty"`

	// **参数解释：** 弹性公网IP的ID。 **取值范围：** 不涉及。
	PublicIpId *string `json:"public_ip_id,omitempty"`

	// **参数解释：** 弹性公网IP对外提供服务的端口号。 **取值范围：** 1~65535。
	ExternalServicePort *int32 `json:"external_service_port,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o BindPublicGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindPublicGatewayResponse struct{}"
	}

	return strings.Join([]string{"BindPublicGatewayResponse", string(data)}, " ")
}
