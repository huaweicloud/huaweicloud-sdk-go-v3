package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkStatisticsResponse Response Object
type ShowNetworkStatisticsResponse struct {

	// 未保护集群数
	ProtectedClusterTotalNum *int32 `json:"protected_cluster_total_num,omitempty"`

	// 集群数
	ClusterTotalNum *int32 `json:"cluster_total_num,omitempty"`

	// 命名空间数
	NamespaceTotalNum *int32 `json:"namespace_total_num,omitempty"`

	// 策略数
	NetworkPolicyTotalNum *int32 `json:"network_policy_total_num,omitempty"`
	HttpStatusCode        int    `json:"-"`
}

func (o ShowNetworkStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowNetworkStatisticsResponse", string(data)}, " ")
}
