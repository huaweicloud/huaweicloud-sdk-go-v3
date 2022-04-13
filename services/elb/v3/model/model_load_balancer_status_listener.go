package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LB状态树的监听器状态信息
type LoadBalancerStatusListener struct {
	// 监听器的名称。

	Name *string `json:"name,omitempty"`
	// 监听器的配置状态。取值： - ACTIVE：使用中。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
	// 监听器下的后端主机组操作状态。

	Pools *[]LoadBalancerStatusPool `json:"pools,omitempty"`
	// 监听器下的7层转发策略操作状态。

	L7policies *[]LoadBalancerStatusPolicy `json:"l7policies,omitempty"`
	// 监听器ID。

	Id *string `json:"id,omitempty"`
	// 监听器的操作状态。取值：  - ONLINE：创建时默认状态，表示监听器正常运行。  - DEGRADED：   -该监听器下存在l7policy或l7rule的Provisioning_status=ERROR时返回这个状态。   -状态树该监听器下存在member的operating_status=OFFLINE。 - DISABLED：负载均衡器或监听器的admin_state_up=false。 使用说明：  - DEGRADED和DISABLED状态仅在当前接口返回，查询监听器详情等其他接口返回字段operating_status不存在这两个状态值。

	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusListener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusListener struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusListener", string(data)}, " ")
}
