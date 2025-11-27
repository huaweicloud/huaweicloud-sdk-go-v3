package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGitopsStatisticsResponse Response Object
type ShowGitopsStatisticsResponse struct {
	OverviewStatus *OverviewReconcileStatus `json:"overviewStatus,omitempty"`

	// 每个集群的配置集合状态
	Items *[]ClusterReconcileStatus `json:"items,omitempty"`

	// 所有页的结果的总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowGitopsStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGitopsStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowGitopsStatisticsResponse", string(data)}, " ")
}
