package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterWithExistedNodesRequestBody 节点扩容请求
type ResizeClusterWithExistedNodesRequestBody struct {
	ScaleOut *ScaleOut `json:"scale_out"`

	Resize *Resize `json:"resize,omitempty"`

	// 是否强制备份
	ForceBackup bool `json:"force_backup"`

	// 扩容模式
	Mode string `json:"mode"`

	// 逻辑集群名称
	LogicalClusterName string `json:"logical_cluster_name"`

	// 是否是使用已添加的空闲节点进行扩容
	ExpandWithExistedNode bool `json:"expand_with_existed_node"`

	// 否只是添加节点
	CreateNodeOnly bool `json:"create_node_only"`

	// 扩容完成后是否自动启动重分布，默认是
	AutoRedistribute bool `json:"auto_redistribute"`

	// 是否调度模式扩容加节点
	IsSchedulerBuildMode bool `json:"is_scheduler_build_mode"`

	RedisConf *RedisConf `json:"redis_conf"`

	BuildTaskInfo *BuildTaskInfo `json:"build_task_info"`

	// 扩容订单ID
	OrderId *string `json:"order_id,omitempty"`
}

func (o ResizeClusterWithExistedNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterWithExistedNodesRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeClusterWithExistedNodesRequestBody", string(data)}, " ")
}
