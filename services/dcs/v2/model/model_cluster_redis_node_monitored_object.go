package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监控对象结构
type ClusterRedisNodeMonitoredObject struct {

	// 测量对象ID，即节点的ID。
	DcsInstanceId *string `json:"dcs_instance_id,omitempty" xml:"dcs_instance_id"`

	// 测量对象名称，即节点IP。
	Name *string `json:"name,omitempty" xml:"name"`

	// 维度dcs_cluster_redis_node的测量对象的ID。
	DcsClusterRedisNode *string `json:"dcs_cluster_redis_node,omitempty" xml:"dcs_cluster_redis_node"`

	// 测量对象状态，即节点状态。
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o ClusterRedisNodeMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterRedisNodeMonitoredObject struct{}"
	}

	return strings.Join([]string{"ClusterRedisNodeMonitoredObject", string(data)}, " ")
}
