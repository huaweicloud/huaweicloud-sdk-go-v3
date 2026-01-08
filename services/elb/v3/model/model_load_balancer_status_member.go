package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatusMember **参数解释**：后端服务器的状态信息。
type LoadBalancerStatusMember struct {

	// **参数解释**：后端服务器配置状态。  **取值范围**：ACTIVE表示使用中。
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	// **参数解释**：后端服务器的IP地址。  **取值范围**：不涉及
	Address *string `json:"address,omitempty"`

	// **参数解释**：后端服务器的端口号。  **取值范围**：1-65535
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// **参数解释**：后端服务器ID。  **取值范围**：不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：后端服务器的操作状态。  **取值范围**： - ONLINE：后端服务器正常运行。 - NO_MONITOR：后端服务器健康检查未开启。 - DISABLED：后端服务器不可用。所属负载均衡器或后端服务器组或该后端服务器的admin_state_up=false时，会出现该状态。注意该状态仅在当前接口中返回。 - OFFLINE：关联ECS已下线。 - INITIAL：后端云服务器健康检查打开时的初始状态。 - UNKNOWN: 后端云服务器组没有绑定监听器或者后端云服务器没有关联ECS等原因，后端云服务器健康检查结果未知。
	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusMember struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusMember", string(data)}, " ")
}
