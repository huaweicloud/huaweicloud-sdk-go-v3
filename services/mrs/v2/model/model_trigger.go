package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Trigger struct {

	// 指标名称。  该触发条件会依据该名称对应指标的值来进行判断。  最大长度为64个字符。
	MetricName string `json:"metric_name" xml:"metric_name"`

	// 指标阈值。  触发该条件的指标阈值，只允许输入整数或者带两位小数的数。
	MetricValue string `json:"metric_value" xml:"metric_value"`

	// 指标判断逻辑运算符，包括：  - LT：小于 - GT：大于 - LTOE：小于等于 - GTOE：大于等于
	ComparisonOperator *string `json:"comparison_operator,omitempty" xml:"comparison_operator"`

	// 判断连续满足指标阈值的周期数(一个周期为5分钟)。  取值范围[1～288]
	EvaluationPeriods int32 `json:"evaluation_periods" xml:"evaluation_periods"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}
