package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Duration 查询时间范围，time_range属性优先于start_time/end_time
type Duration struct {

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 时间范围 - HALF_HOUR: 半小时 - HOUR: 1小时 - THREE_HOUR: 3小时 - TWELVE_HOUR: 12小时 - DAY: 1天 - WEEK: 1周 - MONTH: 1个月
	TimeRange *string `json:"time_range,omitempty"`
}

func (o Duration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Duration struct{}"
	}

	return strings.Join([]string{"Duration", string(data)}, " ")
}
