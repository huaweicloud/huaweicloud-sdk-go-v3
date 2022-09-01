package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListActiveOrHistoryAlarmsRequestBody struct {

	// 关键字检索条件
	Step *int32 `json:"step,omitempty" xml:"step"`

	// 是否自定义查询时间段
	WhetherCustomField bool `json:"whether_custom_field" xml:"whether_custom_field"`

	// 自定义时间段开始时间(时间戳)
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	// 自定义时间段结束时间(时间戳)
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time"`

	// 非自定义时间段时间范围(单位为分钟)
	TimeRange *string `json:"time_range,omitempty" xml:"time_range"`

	// 关键字检索条件
	Search *string `json:"search,omitempty" xml:"search"`

	// 告警级别(\"Critical\",\"Major\",\"Minor\",\"Info\")
	AlarmLevelIds *[]string `json:"alarm_level_ids,omitempty" xml:"alarm_level_ids"`

	Sort *Sort `json:"sort,omitempty" xml:"sort"`
}

func (o ListActiveOrHistoryAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveOrHistoryAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"ListActiveOrHistoryAlarmsRequestBody", string(data)}, " ")
}
