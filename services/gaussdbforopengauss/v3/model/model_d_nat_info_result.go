package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DNatInfoResult struct {

	// **参数解释**: 已经绑定NAT网关的节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: NAT网关实例的ID。 **取值范围**: 不涉及。
	NatGatewayId *string `json:"nat_gateway_id,omitempty"`

	// **参数解释**: 端口ID。 **取值范围**: 不涉及。
	PortId *string `json:"port_id,omitempty"`

	// **参数解释**: 弹性公网ID。 **取值范围**: 不涉及。
	PublicIpId *string `json:"public_ip_id,omitempty"`

	// **参数解释**: 弹性公网IP。 **取值范围**: 不涉及。
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 对外提供服务的端口号，可通过弹性公网IP加该端口号的方式连接数据库实例。 **取值范围**: 不涉及。
	ExternalServicePort *int32 `json:"external_service_port,omitempty"`

	// **参数解释**: GaussDB数据库端口号。 **取值范围**: 不涉及。
	InternalServicePort *int32 `json:"internal_service_port,omitempty"`

	// **参数解释**: 内网地址。 **取值范围**: 不涉及。
	PrivateIp *string `json:"private_ip,omitempty"`
}

func (o DNatInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DNatInfoResult struct{}"
	}

	return strings.Join([]string{"DNatInfoResult", string(data)}, " ")
}
