package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecurrenceSchedule struct {

	// 日期按月
	DayOfMonth *string `json:"day_of_month,omitempty"`

	// 日期按星期
	DayOfWeek *string `json:"day_of_week,omitempty"`

	// 时
	Hour *string `json:"hour,omitempty"`

	// 分
	Minute *string `json:"minute,omitempty"`

	// 月
	Month *string `json:"month,omitempty"`

	// 年
	Year *string `json:"year,omitempty"`
}

func (o RecurrenceSchedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecurrenceSchedule struct{}"
	}

	return strings.Join([]string{"RecurrenceSchedule", string(data)}, " ")
}
