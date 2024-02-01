package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterWithExistedNodesRequestBody 节点扩容请求
type ResizeClusterWithExistedNodesRequestBody struct {
	ScaleOut *ScaleOut `json:"scale_out"`

	// 是否强制备份
	ForceBackup *bool `json:"force_backup,omitempty"`

	// 扩容备份模式，不传默认离线read-only。
	Mode *string `json:"mode,omitempty"`

	// 逻辑集群名称。非逻辑集群模式下该字段不填，逻辑集群模式下不传默认elastic_group。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 是否是使用已添加的空闲节点进行扩容
	ExpandWithExistedNode bool `json:"expand_with_existed_node"`

	// 扩容完成后是否自动启动重分布，默认true
	AutoRedistribute *bool `json:"auto_redistribute,omitempty"`

	RedisConf *RedisConfReq `json:"redis_conf,omitempty"`
}

func (o ResizeClusterWithExistedNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterWithExistedNodesRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeClusterWithExistedNodesRequestBody", string(data)}, " ")
}
