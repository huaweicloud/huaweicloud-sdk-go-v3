package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMissingIndexTrendResponse Response Object
type ShowMissingIndexTrendResponse struct {

	// 趋势数量列表
	TrendList *[]MissingIndexTrendPoint `json:"trend_list,omitempty"`

	UserCostTrend *UserTrendPercent `json:"user_cost_trend,omitempty"`

	UserImpactTrend *UserTrendPercent `json:"user_impact_trend,omitempty"`

	UserSeekTrend  *UserSeekTrend `json:"user_seek_trend,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowMissingIndexTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMissingIndexTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowMissingIndexTrendResponse", string(data)}, " ")
}
