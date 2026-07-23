package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScheduleInfo 创建策略信息。
type JobScheduleInfo struct {

	// 策略id。
	Id string `json:"id"`

	// 是否是用户自定义模板。false表示为系统默认模板。true表示为用户自定义模板。
	UserDefined bool `json:"user_defined"`

	// 策略类型，snapshot:快照策略, sync:同步策略。
	ScheduleType string `json:"schedule_type"`

	// 计划类型。  automatically：SQL Server代理启动时自动启动。 cpu_idle：CPU空闲时启动。 recurring：重复执行。 one_time：执行一次。
	JobScheduleType *string `json:"job_schedule_type,omitempty"`

	OneTimeOccurrence *JobScheduleOneTimeOccurrenceInfo `json:"one_time_occurrence,omitempty"`

	Frequency *JobScheduleFrequencyInfo `json:"frequency"`

	DailyFrequency *JobScheduleDailyFrequencyInfo `json:"daily_frequency"`

	Duration *JobScheduleDurationInfo `json:"duration"`
}

func (o JobScheduleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScheduleInfo struct{}"
	}

	return strings.Join([]string{"JobScheduleInfo", string(data)}, " ")
}
