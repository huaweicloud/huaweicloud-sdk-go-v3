package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatusListener **参数解释**：监听器的状态信息。
type LoadBalancerStatusListener struct {

	// **参数解释**：监听器的名称。  **取值范围**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：监听器的配置状态。  **取值范围**： - ACTIVE：使用中。
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	// **参数解释**：监听器下的所有后端服务器组的状态信息。
	Pools *[]LoadBalancerStatusPool `json:"pools,omitempty"`

	// **参数解释**：监听器下的7层转发策略的状态信息。
	L7policies *[]LoadBalancerStatusPolicy `json:"l7policies,omitempty"`

	// **参数解释**：监听器ID。  **取值范围**：不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：监听器的操作状态。  **取值范围**： - ONLINE：创建时默认状态，表示监听器正常运行。 - DEGRADED：该监听器下存在l7policy或l7rule的Provisioning_status=ERROR时返回这个状态。或者状态树该监听器下存在member的operating_status=OFFLINE。 - DISABLED：负载均衡器或监听器的admin_state_up=false。  > DEGRADED和DISABLED状态仅在当前接口返回，查询监听器详情等其他接口返回字段operating_status不存在这两个状态值。
	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusListener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusListener struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusListener", string(data)}, " ")
}
