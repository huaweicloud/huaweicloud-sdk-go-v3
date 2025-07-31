package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowContainerNodeStatisticsResponse Response Object
type ShowContainerNodeStatisticsResponse struct {

	// 未防护服务器数
	UnprotectedNum *int32 `json:"unprotected_num,omitempty"`

	// 开启防护节点
	ProtectedNum *int32 `json:"protected_num,omitempty"`

	// 按需开启防护节点
	ProtectedNumOnDemand *int32 `json:"protected_num_on_demand,omitempty"`

	// 按配额开启防护节点
	ProtectedNumPacketCycle *int32 `json:"protected_num_packet_cycle,omitempty"`

	// 未安装集群节点
	ClusterNodeNotInstalledNum *int32 `json:"cluster_node_not_installed_num,omitempty"`

	// 未安装非集群节点
	NotClusterNodeNotInstalledNum *int32 `json:"not_cluster_node_not_installed_num,omitempty"`
	HttpStatusCode                int    `json:"-"`
}

func (o ShowContainerNodeStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowContainerNodeStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowContainerNodeStatisticsResponse", string(data)}, " ")
}
