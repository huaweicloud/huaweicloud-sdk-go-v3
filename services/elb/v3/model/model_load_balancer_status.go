package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatus **参数解释**：负载均衡器及其子资源的状态信息。
type LoadBalancerStatus struct {

	// **参数解释**：负载均衡器名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：负载均衡器的配置状态。  **取值范围**： - ACTIVE：使用中。 - PENDING_DELETE：删除中。
	ProvisioningStatus string `json:"provisioning_status"`

	// **参数解释**：负载均衡器关联的所有监听器的状态信息。
	Listeners []LoadBalancerStatusListener `json:"listeners"`

	// **参数解释**：负载均衡器关联的所有后端服务器组的状态信息。
	Pools []LoadBalancerStatusPool `json:"pools"`

	// **参数解释**：负载均衡器ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：负载均衡器的操作状态。  **取值范围**： - ONLINE：创建时默认状态，表示负载均衡器正常运行。 - FROZEN：已冻结。 - DEGRADED：负载均衡器下存在member的operating_status为OFFLINE时返回这个状态。 - DISABLED：负载均衡器的admin_state_up属性值为false。  > DEGRADED和DISABLED状态仅在当前接口中返回，查询负载均衡器详情等其他接口不会返回这两个状态值。
	OperatingStatus string `json:"operating_status"`
}

func (o LoadBalancerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatus struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatus", string(data)}, " ")
}
