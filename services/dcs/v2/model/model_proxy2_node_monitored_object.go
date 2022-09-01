package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Redis 4.0 和 5.0 proxy集群中proxy节点监控对象结构
type Proxy2NodeMonitoredObject struct {

	// 测量对象ID，即节点的ID。
	DcsInstanceId *string `json:"dcs_instance_id,omitempty" xml:"dcs_instance_id"`

	// 测量对象名称，即节点IP。
	Name *string `json:"name,omitempty" xml:"name"`

	// 维度dcs_cluster_proxy2_node 的测量对象的ID。
	DcsClusterProxy2Node *string `json:"dcs_cluster_proxy2_node,omitempty" xml:"dcs_cluster_proxy2_node"`

	// 测量对象状态，即节点状态。
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o Proxy2NodeMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Proxy2NodeMonitoredObject struct{}"
	}

	return strings.Join([]string{"Proxy2NodeMonitoredObject", string(data)}, " ")
}
