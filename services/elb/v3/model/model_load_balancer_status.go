package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoadBalancerStatus struct {

	// 负载均衡器名称。
	Name string `json:"name"`

	// 负载均衡器的配置状态。  取值： - ACTIVE：使用中。 - PENDING_DELETE：删除中。
	ProvisioningStatus string `json:"provisioning_status"`

	// 负载均衡器关联的监听器列表。
	Listeners []LoadBalancerStatusListener `json:"listeners"`

	// 负载均衡器关联的后端服务器组列表。
	Pools []LoadBalancerStatusPool `json:"pools"`

	// 负载均衡器ID。
	Id string `json:"id"`

	// 负载均衡器的操作状态。  取值： - ONLINE：创建时默认状态，表示负载均衡器正常运行。 - FROZEN：已冻结。 - DEGRADED：负载均衡器下存在member的operating_status为OFFLINE时返回这个状态。 - DISABLED：负载均衡器的admin_state_up属性值为false。  说明：DEGRADED和DISABLED状态仅在当前接口中返回，LB详情等其他接口不返回这两个状态值。
	OperatingStatus string `json:"operating_status"`
}

func (o LoadBalancerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatus struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatus", string(data)}, " ")
}
