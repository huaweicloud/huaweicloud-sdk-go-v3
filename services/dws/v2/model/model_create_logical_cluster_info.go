package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogicalClusterInfo 创建逻辑集群请求信息
type CreateLogicalClusterInfo struct {

	// 逻辑集群名称
	LogicalClusterName string `json:"logical_cluster_name"`

	// 逻辑集群环信息
	ClusterRings []ClusterRing `json:"cluster_rings"`
}

func (o CreateLogicalClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogicalClusterInfo struct{}"
	}

	return strings.Join([]string{"CreateLogicalClusterInfo", string(data)}, " ")
}
