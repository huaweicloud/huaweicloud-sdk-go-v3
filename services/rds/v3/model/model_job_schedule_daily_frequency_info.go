package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScheduleDailyFrequencyInfo 策略每日频率。
type JobScheduleDailyFrequencyInfo struct {

	// 每日频率类型 once:每日一次, multiple:每日多次。
	FreqSubdayType *string `json:"freq_subday_type,omitempty"`

	// 每日第一次执行时间。每日频率类型为每日一次时，则只执行这一次。 HH:mm:ss
	ActiveStartTime *string `json:"active_start_time,omitempty"`

	// 最后一次执行时间。每日执行多次时该参数必传，每日执行一次时不生效。 HH:mm:ss
	ActiveEndTime *string `json:"active_end_time,omitempty"`

	// 执行间隔。每日执行多次时该参数必传，每日执行一次时不生效。取值范围1至99
	FreqSubdayInterval *int32 `json:"freq_subday_interval,omitempty"`

	// 执行间隔的单位。每日执行多次时该参数必传，每日执行一次时不生效。 second:秒, minute:分, hour:时
	FreqIntervalUnit *string `json:"freq_interval_unit,omitempty"`
}

func (o JobScheduleDailyFrequencyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScheduleDailyFrequencyInfo struct{}"
	}

	return strings.Join([]string{"JobScheduleDailyFrequencyInfo", string(data)}, " ")
}
