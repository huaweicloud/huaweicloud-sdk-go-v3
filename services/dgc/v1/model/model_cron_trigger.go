package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CronTrigger struct {
	StartTime string `json:"startTime"`

	EndTime *string `json:"endTime,omitempty"`

	// Cron表达式
	Expression string `json:"expression"`

	// 时区
	ExpressionTimeZone *string `json:"expressionTimeZone,omitempty"`

	Period string `json:"period"`

	// 是否依赖本作业上一个运行周期任务的执行结果
	DependPrePeriod *bool `json:"dependPrePeriod,omitempty"`

	// 依赖其它作业列表
	DependJobs *[]DependJobs `json:"dependJobs,omitempty"`

	// 并发调用数
	Concurrent *int32 `json:"concurrent,omitempty"`
}

func (o CronTrigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CronTrigger struct{}"
	}

	return strings.Join([]string{"CronTrigger", string(data)}, " ")
}
