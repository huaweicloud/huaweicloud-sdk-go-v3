package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterAssetStatisticsResponse Response Object
type ShowClusterAssetStatisticsResponse struct {

	// 集群数量
	ClusterNum *int32 `json:"cluster_num,omitempty"`

	// 工作负载数量
	WorkLoadNum *int32 `json:"work_load_num,omitempty"`

	// 服务数量
	ServiceNum *int32 `json:"service_num,omitempty"`

	// pod数量
	PodNum         *int32 `json:"pod_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowClusterAssetStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterAssetStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterAssetStatisticsResponse", string(data)}, " ")
}
