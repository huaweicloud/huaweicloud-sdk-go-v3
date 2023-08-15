package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RankEtlFilter
type RankEtlFilter struct {

	// 行为去重方式： - abs_weight，权重绝对值 - date，日期
	FilterType string `json:"filter_type"`

	// 时间类型： - day，天 - week，周 - month，月
	TimeType string `json:"time_type"`

	// 周一是否是第一天。
	IsMondayFirst *bool `json:"is_monday_first,omitempty"`
}

func (o RankEtlFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RankEtlFilter struct{}"
	}

	return strings.Join([]string{"RankEtlFilter", string(data)}, " ")
}
