package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterScanStatisticsResponse Response Object
type ShowClusterScanStatisticsResponse struct {

	// 有风险集群数量
	RiskClusterNum *int32 `json:"risk_cluster_num,omitempty"`

	// 存在应用漏洞的集群数量
	AppVulClusterNum *int32 `json:"app_vul_cluster_num,omitempty"`

	// 未扫描集群数量
	UnscanClusterNum *int32 `json:"unscan_cluster_num,omitempty"`

	// 集群总数量
	AllClusterNum  *int32 `json:"all_cluster_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowClusterScanStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterScanStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterScanStatisticsResponse", string(data)}, " ")
}
