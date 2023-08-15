package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SchedulerInfo struct {

	// 表达式
	CronExpression *string `json:"cron_expression,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 最大超时时间
	MaxTimeOut *int32 `json:"max_time_out,omitempty"`

	// 间隔
	Interval *string `json:"interval,omitempty"`

	// 调度类型
	ScheduleType *string `json:"schedule_type,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 工作id
	JobId *int32 `json:"job_id,omitempty"`

	// 启用
	Enabled *int32 `json:"enabled,omitempty"`
}

func (o SchedulerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchedulerInfo struct{}"
	}

	return strings.Join([]string{"SchedulerInfo", string(data)}, " ")
}
