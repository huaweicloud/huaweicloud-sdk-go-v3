package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitoredObjectsOfInstanceResponse Response Object
type ListMonitoredObjectsOfInstanceResponse struct {

	// 当前查询维度路由。如果是主维度，则数组中是自身ID。
	Router *[]string `json:"router,omitempty"`

	// 当前查询维度子维度对象列表。当前只有维度为dcs_instance_id时才有值。 - Proxy集群有两个子维度，分别为dcs_cluster_redis_node和dcs_cluster_proxy_node。 - Cluster集群有一个子维度 dcs_cluster_proxy_node。
	Children *[]DimChild `json:"children,omitempty"`

	// 当前查询维度监控对象列表。
	Instances *[]InstancesMonitoredObject `json:"instances,omitempty"`

	// Proxy集群或Cluster集群时才存在，表示集群数据节点维度的监控对象列表。字段名称与children的子维度对象名称相同。
	DcsClusterRedisNode *[]ClusterRedisNodeMonitoredObject `json:"dcs_cluster_redis_node,omitempty"`

	// Proxy集群时才存在，表示集群Proxy节点维度的监控对象列表。字段名称与children的子维度对象名称相同。
	DcsClusterProxyNode *[]ProxyNodeMonitoredObject `json:"dcs_cluster_proxy_node,omitempty"`

	// Redis 4.0和5.0的Proxy集群时才存在，表示集群Proxy节点维度的监控对象列表。字段名称与children的子维度对象名称相同。
	DcsClusterProxy2Node *[]Proxy2NodeMonitoredObject `json:"dcs_cluster_proxy2_node,omitempty"`

	// 主维度监控对象的总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMonitoredObjectsOfInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitoredObjectsOfInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListMonitoredObjectsOfInstanceResponse", string(data)}, " ")
}
