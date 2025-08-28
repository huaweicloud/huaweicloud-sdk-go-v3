package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadBalancerPortsRequest Request Object
type ShowLoadBalancerPortsRequest struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	// **参数解释**：负载均衡器占用的端口ID。  支持多值查询，查询条件格式：*port_id=xxx&port_id=xxx*。
	PortId *[]string `json:"port_id,omitempty"`

	// **参数解释**：负载均衡器占用的私有IPv4地址。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及  支持多值查询，查询条件格式：*ip_address=xxx&ip_address=xxx*。
	IpAddress *[]string `json:"ip_address,omitempty"`

	// **参数解释**：负载均衡器占用的IPv6地址。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及  支持多值查询，查询条件格式：*ipv6_address=xxx&ipv6_address=xxx*。
	Ipv6Address *[]string `json:"ipv6_address,omitempty"`

	// **参数解释**：子网端口类型。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及  支持多值查询，查询条件格式：*type=xxx&type=xxx*。
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：子网端口所在下联面子网网络ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及  支持多值查询，查询条件格式：*virsubnet_id=xxx&virsubnet_id=xxx*。
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`
}

func (o ShowLoadBalancerPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerPortsRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerPortsRequest", string(data)}, " ")
}
