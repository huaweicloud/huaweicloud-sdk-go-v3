package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceClusterInfo 使用已有CCE集群信息，说明：Fabric1.4版本服务仅支持1.15及以下版本集群
type CceClusterInfo struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 集群CPU架构类型：X86（VirtualMachine），ARM（ARM64）
	ClusterPlatformType string `json:"cluster_platform_type"`
}

func (o CceClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceClusterInfo struct{}"
	}

	return strings.Join([]string{"CceClusterInfo", string(data)}, " ")
}
