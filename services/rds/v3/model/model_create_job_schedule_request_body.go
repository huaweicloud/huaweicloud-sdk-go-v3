package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobScheduleRequestBody 创建策略信息。
type CreateJobScheduleRequestBody struct {

	// 策略类型，snapshot:快照策略, sync:同步策略。
	ScheduleType string `json:"schedule_type"`

	// 计划类型。默认值recurring。  automatically：SQL Server代理启动时自动启动。 cpu_idle：CPU空闲时启动。 recurring：重复执行。 one_time：执行一次。
	JobScheduleType *string `json:"job_schedule_type,omitempty"`

	OneTimeOccurrence *JobScheduleOneTimeOccurrenceInfo `json:"one_time_occurrence,omitempty"`

	Frequency *JobScheduleFrequencyInfo `json:"frequency,omitempty"`

	DailyFrequency *JobScheduleDailyFrequencyInfo `json:"daily_frequency,omitempty"`

	Duration *JobScheduleDurationInfo `json:"duration,omitempty"`
}

func (o CreateJobScheduleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobScheduleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateJobScheduleRequestBody", string(data)}, " ")
}
