package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Cron 当type为CRON时，配置调度频率、启动时间等信息
type Cron struct {

	// 调度开始时间，采用ISO 8601时间表示方法，格式为yyyy-MM-dd'T'HH:mm:ssZ，例如2018-10-22T23:59:59+08表示的时间为2018年10月22日23时59分59秒，在正8区，即北京时区。
	StartTime string `json:"start_time"`

	// 调度结束时间，采用ISO 8601时间表示方法，格式为yyyy-MM-dd'T'HH:mm:ssZ，例如2018-10-22T23:59:59+08表示的时间为2018年10月22日23时59分59秒，在正8区，即北京时区。如果结束时间不配置，作业会按照调度周期一直执行下去。
	EndTime *string `json:"end_time,omitempty"`

	// Cron表达式，格式为\"<秒> <分> <时> <天> <月> <星期>\"
	Expression string `json:"expression"`

	// Cron表达式对应的时区信息，例如GMT+8。默认值：使用DataArts Studio服务端所在的时区。
	ExpressionTimeZone *string `json:"expression_time_zone,omitempty"`

	// 用于描述运行的间隔时间，格式为时间+时间单位。需要与expression中的表达式对应
	Period string `json:"period"`

	// 是否依赖本作业上一个运行周期任务的执行结果
	DependPrePeriod *bool `json:"depend_pre_period,omitempty"`

	// 依赖其它作业列表
	DependJobs *[]DependJob `json:"depend_jobs,omitempty"`
}

func (o Cron) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cron struct{}"
	}

	return strings.Join([]string{"Cron", string(data)}, " ")
}
