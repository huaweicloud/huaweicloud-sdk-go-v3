package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CronTrigger struct {

	// 调度开始时间，采用ISO 8601时间表示方法，格式为yyyy-MM-dd'T'HH:mm:ssZ，例如2018-10-22T23:59:59+08表示的时间为2018年10月22日23时59分59秒，在正8区，即北京时区。
	StartTime string `json:"startTime"`

	// 调度结束时间，采用ISO 8601时间表示方法，格式为yyyy-MM-dd'T'HH:mm:ssZ，例如2018-10-22T23:59:59+08表示的时间为2018年10月22日23时59分59秒，在正8区，即北京时区。如果结束时间不配置，作业会按照调度周期一直执行下去。
	EndTime *string `json:"endTime,omitempty"`

	// Cron表达式
	Expression string `json:"expression"`

	// 时区
	ExpressionTimeZone *string `json:"expressionTimeZone,omitempty"`

	// 用于描述运行的间隔时间，格式为时间+时间单位。需要与expression中的表达式对应
	Period string `json:"period"`

	// 是否依赖本作业上一个运行周期任务的执行结果
	DependPrePeriod *bool `json:"dependPrePeriod,omitempty"`

	DependJobs *DependJob `json:"dependJobs,omitempty"`

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
