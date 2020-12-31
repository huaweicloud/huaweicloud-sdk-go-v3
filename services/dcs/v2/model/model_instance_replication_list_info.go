/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 副本列表
type InstanceReplicationListInfo struct {
	// 副本角色，取值有： - master：表示主节点。 - slave：表示从节点。
	ReplicationRole *string `json:"replication_role,omitempty"`
	// 副本IP。
	ReplicationIp *string `json:"replication_ip,omitempty"`
	// 是否是新加副本。
	IsReplication *bool `json:"is_replication,omitempty"`
	// 副本id。
	ReplicationId *string `json:"replication_id,omitempty"`
	// 节点id。
	NodeId *string `json:"node_id,omitempty"`
	// 副本状态。
	Status *string `json:"status,omitempty"`
	// 副本所在的可用区
	AzCode *string `json:"az_code,omitempty"`
	// 副本对应的监控指标维度信息。可用于调用云监控服务的查询监控数据指标相关接口 - 副本的监控维度为多维度，返回数组中包含两个维度信息。从云监控查询监控数据时，要按多维度传递指标维度参数，才能查询到监控指标值 - 第一个维度为副本父维度信息，维度名称为dcs_instance_id，维度值对应副本所在的实例ID - 第二个维度，维度名称为dcs_cluster_redis_node,维度值为副本的监控对象ID，与副本ID和节点ID不同。
	Dimensions *[]InstanceReplicationDimensionsInfo `json:"dimensions,omitempty"`
}

func (o InstanceReplicationListInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InstanceReplicationListInfo struct{}"
	}

	return strings.Join([]string{"InstanceReplicationListInfo", string(data)}, " ")
}
