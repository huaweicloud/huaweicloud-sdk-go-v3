package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Redis 3.0 proxy集群中proxy节点监控对象结构
type ProxyNodeMonitoredObject struct {
	// 测量对象ID，即节点的ID。

	DcsInstanceId *string `json:"dcs_instance_id,omitempty"`
	// 测量对象名称，即节点IP。

	Name *string `json:"name,omitempty"`
	// 维度dcs_cluster_proxy_node 的测量对象的ID。

	DcsClusterProxyNode *string `json:"dcs_cluster_proxy_node,omitempty"`
	// 测量对象状态，即节点状态。

	Status *string `json:"status,omitempty"`
}

func (o ProxyNodeMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyNodeMonitoredObject struct{}"
	}

	return strings.Join([]string{"ProxyNodeMonitoredObject", string(data)}, " ")
}
