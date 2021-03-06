package model

import (
	"encoding/json"

	"strings"
)

// lb状态树的监听器状态信息
type LoadBalancerStatusListener struct {
	// 负载均衡器下监听器的名称。

	Name *string `json:"name,omitempty"`
	// provisioning状态。 可以为ACTIVE、PENDING_CREATE 或者ERROR。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
	// 挂载在监听器下的后端主机组。

	Pools *[]LoadBalancerStatusPool `json:"pools,omitempty"`
	// 7层转发策略

	L7policies *[]LoadBalancerStatusPolicy `json:"l7policies,omitempty"`
	// 监听器ID。

	Id *string `json:"id,omitempty"`
	// 操作状态。 可以为：ONLINE、OFFLINE、DEGRADED、DISABLED或NO_MONITOR。说明：该字段为预留字段，暂未启用，默认为ONLINE。

	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusListener) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusListener struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusListener", string(data)}, " ")
}
