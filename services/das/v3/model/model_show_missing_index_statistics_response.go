package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMissingIndexStatisticsResponse Response Object
type ShowMissingIndexStatisticsResponse struct {

	// 采集时间（ms）
	CollectTime *int64 `json:"collect_time,omitempty"`

	// 索引缺失总数
	TotalMissingIndexCount *int64 `json:"total_missing_index_count,omitempty"`

	// 性能提示大于80%的数量
	UserImpactGt80Count *int64 `json:"user_impact_gt80_count,omitempty"`

	// 近1天用户访问条数
	LastDayAccessedCount *int64 `json:"last_day_accessed_count,omitempty"`

	// 近1周用户访问条数
	LastWeekAccessedCount *int64 `json:"last_week_accessed_count,omitempty"`

	// 近2周用户访问条数
	LastTwoWeekAccessedCount *int64 `json:"last_two_week_accessed_count,omitempty"`

	// 近1月用户访问条数
	LastMonthAccessedCount *int64 `json:"last_month_accessed_count,omitempty"`
	HttpStatusCode         int    `json:"-"`
}

func (o ShowMissingIndexStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMissingIndexStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowMissingIndexStatisticsResponse", string(data)}, " ")
}
