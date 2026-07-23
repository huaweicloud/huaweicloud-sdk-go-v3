package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyJobScheduleRequestBody 修改策略信息。
type ModifyJobScheduleRequestBody struct {

	// 计划类型。默认值recurring。  automatically：SQL Server代理启动时自动启动。 cpu_idle：CPU空闲时启动。 recurring：重复执行。 one_time：执行一次。
	JobScheduleType *string `json:"job_schedule_type,omitempty"`

	OneTimeOccurrence *JobScheduleOneTimeOccurrenceInfo `json:"one_time_occurrence,omitempty"`

	Frequency *JobScheduleFrequencyInfo `json:"frequency,omitempty"`

	DailyFrequency *JobScheduleDailyFrequencyInfo `json:"daily_frequency,omitempty"`

	Duration *JobScheduleDurationInfo `json:"duration,omitempty"`
}

func (o ModifyJobScheduleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobScheduleRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyJobScheduleRequestBody", string(data)}, " ")
}
