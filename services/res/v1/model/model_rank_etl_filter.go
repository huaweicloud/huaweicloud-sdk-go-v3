package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RankEtlFilter struct {

	// 行为去重方式： - abs_weight，权重绝对值 - date，日期
	FilterType string `json:"filter_type" xml:"filter_type"`

	// 时间类型： - day，天 - week，周 - month，月
	TimeType string `json:"time_type" xml:"time_type"`

	// 周一是否是第一天。
	IsMondayFirst *bool `json:"is_monday_first,omitempty" xml:"is_monday_first"`
}

func (o RankEtlFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RankEtlFilter struct{}"
	}

	return strings.Join([]string{"RankEtlFilter", string(data)}, " ")
}
