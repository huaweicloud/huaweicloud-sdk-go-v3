package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Cron struct {

	// 调度开始时间，采用ISO 8601时间表示方法，格式为yyyy-MM-dd'T'HH:mm:ssZ，例如2018-10-22T23:59:59+08表示的时间为2018年10月22日23时59分59秒，在正8区，即北京时区
	StartTime string `json:"startTime"`

	// 调度结束时间，采用ISO 8601时间表示方法，格式为yyyy-MM-dd'T'HH:mm:ssZ，例如2018-10-22T23:59:59+08表示的时间为2018年10月22日23时59分59秒，在正8区，即北京时区。如果结束时间不配置，作业会按照调度周期一直执行下去
	EndTime *string `json:"endTime,omitempty"`

	// Cron表达式，格式为\"<秒> <分> <时> <天> <月> <星期>\"
	Expression string `json:"expression"`

	// Cron表达式对应的时区信息，例如GMT+8
	ExpressionTimeZone *string `json:"expressionTimeZone,omitempty"`

	// 是否依赖本作业上一个运行周期任务的执行结果
	DependPrePeriod *bool `json:"dependPrePeriod,omitempty"`

	// 调度间隔类型： - minutes：分钟 - hours：小时 - days：天 - weeks： 周 - months：月 - seasons： 季 - years：年
	IntervalType *string `json:"intervalType,omitempty"`

	DependJobs *DependJob `json:"dependJobs,omitempty"`
}

func (o Cron) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cron struct{}"
	}

	return strings.Join([]string{"Cron", string(data)}, " ")
}
