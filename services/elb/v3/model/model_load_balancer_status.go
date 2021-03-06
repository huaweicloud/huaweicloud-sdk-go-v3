package model

import (
	"encoding/json"

	"strings"
)

//
type LoadBalancerStatus struct {
	// 负载均衡器名称。

	Name string `json:"name"`
	// 负载均衡器的配置状态。 可以为：ACTIVE、PENDING_CREATE 或者ERROR。说明：该字段为预留字段，暂未启用，默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`
	// 负载均衡器关联的监听器列表。

	Listeners []LoadBalancerStatusListener `json:"listeners"`
	// 负载均衡器关联的后端云服务器组列表。

	Pools []LoadBalancerStatusPool `json:"pools"`
	// 负载均衡器ID

	Id string `json:"id"`
	// 负载均衡器的操作状态。 可以为：ONLINE、OFFLINE、DEGRADED、DISABLED或NO_MONITOR。说明：该字段为预留字段，暂未启用，默认为ONLINE。

	OperatingStatus string `json:"operating_status"`
}

func (o LoadBalancerStatus) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatus struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatus", string(data)}, " ")
}
