package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListActiveOrHistoryAlarmsRequestBody struct {
	// 关键字检索条件

	Step *int32 `json:"step,omitempty"`
	// 是否自定义查询时间段

	WhetherCustomField bool `json:"whether_custom_field"`
	// 自定义时间段开始时间(时间戳)

	StartTime *int64 `json:"start_time,omitempty"`
	// 自定义时间段结束时间(时间戳)

	EndTime *int64 `json:"end_time,omitempty"`
	// 非自定义时间段时间范围(单位为分钟)

	TimeRange *string `json:"time_range,omitempty"`
	// 关键字检索条件

	Search *string `json:"search,omitempty"`
	// 告警级别(\"Critical\",\"Major\",\"Minor\",\"Info\")

	AlarmLevelIds *[]string `json:"alarm_level_ids,omitempty"`

	Sort *Sort `json:"sort,omitempty"`
}

func (o ListActiveOrHistoryAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveOrHistoryAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"ListActiveOrHistoryAlarmsRequestBody", string(data)}, " ")
}
