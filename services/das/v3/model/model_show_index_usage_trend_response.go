package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndexUsageTrendResponse Response Object
type ShowIndexUsageTrendResponse struct {

	// 趋势数量列表
	TrendList *[]IndexUsageTrendPoint `json:"trend_list,omitempty"`

	FragmentationTrend *IndexUsagePercent `json:"fragmentation_trend,omitempty"`

	UsageTrend     *IndexUsagePercent `json:"usage_trend,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowIndexUsageTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndexUsageTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowIndexUsageTrendResponse", string(data)}, " ")
}
