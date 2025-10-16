package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScheduleFrequencyInfo 策略间隔周期。
type JobScheduleFrequencyInfo struct {

	// 策略频率类型 daily:按天, weekly:按周, monthly_day:按月、每月按天, monthly_week:按月、每月按周。
	FreqType *string `json:"freq_type,omitempty"`

	// 执行间隔。取值范围1至99。
	FreqInterval *int32 `json:"freq_interval,omitempty"`

	// 频率类型为按周时该参数必传，不为按周时不生效。 每周执行哪几天。参数值：Monday，Tuesday … Sunday
	FreqIntervalWeekly *[]string `json:"freq_interval_weekly,omitempty"`

	// 频率类型为按月、每月按天时该参数必传，不为按月、每月按天时时不生效。 每月执行的日期。取值范围1-31。
	FreqIntervalDayMonthly *int32 `json:"freq_interval_day_monthly,omitempty"`

	// 频率类型为按月、每月按周时该参数必传，不为按月、每月按周时时不生效。 每周执行哪几天。 Sunday, Monday,Tuesday ... Saturday, day, weekday, weekend
	FreqIntervalMonthly *string `json:"freq_interval_monthly,omitempty"`

	// 频率类型为按月、每月按周时该参数必传，不为按月、每月按周时时不生效。 每月在哪周执行。 first, second, third, fourth, last
	FreqRelativeIntervalMonthly *string `json:"freq_relative_interval_monthly,omitempty"`
}

func (o JobScheduleFrequencyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScheduleFrequencyInfo struct{}"
	}

	return strings.Join([]string{"JobScheduleFrequencyInfo", string(data)}, " ")
}
