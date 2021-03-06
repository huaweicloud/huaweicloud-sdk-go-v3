package model

import (
	"encoding/json"

	"strings"
)

// lb状态树的主机组状态信息
type LoadBalancerStatusPool struct {
	// provisioning的状态。 可以为：ACTIVE、PENDING_CREATE 或者ERROR。说明：该字段为预留字段，暂未启用，默认为ACTIVE。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
	// 后端服务器组名。

	Name *string `json:"name,omitempty"`

	Healthmonitor *LoadBalancerStatusHealthMonitor `json:"healthmonitor,omitempty"`
	// 后端服务器。

	Members *[]LoadBalancerStatusMember `json:"members,omitempty"`
	// 后端服务器组ID。

	Id *string `json:"id,omitempty"`
	// 操作状态。 可以为：ONLINE、OFFLINE、DEGRADED、DISABLED或NO_MONITOR。说明：该字段为预留字段，暂未启用，默认为ONLINE。

	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusPool) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPool struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPool", string(data)}, " ")
}
