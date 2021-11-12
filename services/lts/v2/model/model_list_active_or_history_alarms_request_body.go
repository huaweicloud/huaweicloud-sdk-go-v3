package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListActiveOrHistoryAlarmsRequestBody struct {
	// 关键字检索条件

	Step *int32 `json:"step,omitempty"`
	// 是否自定义查询时间段

	WhetherCustomField bool `json:"whether_custom_field"`
	// 自定义时间段开始时间(时间戳)

	StartTime *interface{} `json:"start_time,omitempty"`
	// 自定义时间段结束时间(时间戳)

	EndTime *interface{} `json:"end_time,omitempty"`
	// 非自定义时间段时间范围(单位为分钟)

	TimeRange *int32 `json:"time_range,omitempty"`
	// 关键字检索条件

	Search *string `json:"search,omitempty"`
	// 告警级别(\"Critical\",\"Major\",\"Minor\",\"Info\")

	AlarmLevelIds *[]string `json:"alarmLevelIds,omitempty"`
	// 排序检索条件

	Sort *Sort `json:"sort,omitempty"`
}

func (o ListActiveOrHistoryAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveOrHistoryAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"ListActiveOrHistoryAlarmsRequestBody", string(data)}, " ")
}
