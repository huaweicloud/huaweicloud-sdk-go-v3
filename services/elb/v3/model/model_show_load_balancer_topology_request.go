package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadBalancerTopologyRequest Request Object
type ShowLoadBalancerTopologyRequest struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	// 监听器的ID。  支持多值查询，查询条件格式：*listener_id=xxx&listener_id=xxx*。
	ListenerId *string `json:"listener_id,omitempty"`

	// 后端服务器组的ID。  支持多值查询，查询条件格式：*pool_id=xxx&pool_id=xxx*。
	PoolId *string `json:"pool_id,omitempty"`

	// 监听器的名称。  支持多值查询，查询条件格式：*listener_name=xxx&listener_name=xxx*。
	ListenerName *string `json:"listener_name,omitempty"`

	// 监听器的协议。  支持多值查询，查询条件格式：*listener_protocol=xxx&listener_protocol=xxx*。
	ListenerProtocol *string `json:"listener_protocol,omitempty"`

	// 监听器的监听端口。  支持多值查询，查询条件格式：*listener_protocol_port=xxx&listener_protocol_port=xxx*。
	ListenerProtocolPort *int32 `json:"listener_protocol_port,omitempty"`

	// 后端服务器组的名称。  支持多值查询，查询条件格式：*pool_name=xxx&pool_name=xxx*。
	PoolName *string `json:"pool_name,omitempty"`
}

func (o ShowLoadBalancerTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerTopologyRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerTopologyRequest", string(data)}, " ")
}
