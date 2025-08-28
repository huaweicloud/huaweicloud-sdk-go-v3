package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocalPort **参数解释**：当前ELB占用的子网端口。  **取值范围**：不涉及
type LocalPort struct {

	// **参数解释**：负载均衡器占用的端口ID。  **取值范围**：不涉及
	PortId *string `json:"port_id,omitempty"`

	// **参数解释**：负载均衡器占用的私有IPv4地址。  **取值范围**：不涉及
	IpAddress *string `json:"ip_address,omitempty"`

	// **参数解释**：负载均衡器占用的IPv6地址。  **取值范围**：不涉及
	Ipv6Address *string `json:"ipv6_address,omitempty"`

	// **参数解释**：子网端口类型。  **取值范围**： - l4_hc：四层dnat转发时健康检查使用的地址。 - l4 四层fullnat 转发时健康检查及业务转发使用的地址。 - l7 七层健康检查及业务转发使用的地址。
	Type *string `json:"type,omitempty"`

	// **参数解释**：子网端口所在下联面子网网络ID。  **取值范围**：不涉及
	VirsubnetId *string `json:"virsubnet_id,omitempty"`
}

func (o LocalPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalPort struct{}"
	}

	return strings.Join([]string{"LocalPort", string(data)}, " ")
}
