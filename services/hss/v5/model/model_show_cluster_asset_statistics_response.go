package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterAssetStatisticsResponse Response Object
type ShowClusterAssetStatisticsResponse struct {

	// **参数解释**: 集群数量 **取值范围**: 最小值0，最大值65535
	ClusterNum *int32 `json:"cluster_num,omitempty"`

	// **参数解释**: 工作负载数量 **取值范围**: 最小值0，最大值65535
	WorkLoadNum *int32 `json:"work_load_num,omitempty"`

	// **参数解释**: 服务数量 **取值范围**: 最小值0，最大值65535
	ServiceNum *int32 `json:"service_num,omitempty"`

	// **参数解释**: pod数量 **取值范围**: 最小值0，最大值65535
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
