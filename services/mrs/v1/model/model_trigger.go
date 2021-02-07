package model

import (
	"encoding/json"

	"strings"
)

type Trigger struct {
	// 指标名称。  该触发条件会依据该名称对应指标的值来进行判断。  最大长度为64个字符。  支持的指标名称请参考[表12](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__t27de3279a99a48968dacb015c498d9cb)。
	MetricName string `json:"metric_name"`
	// 指标阈值。  触发该条件的指标阈值，只允许输入整数或者带两位小数的数，metric_name对应的指标数值类型和有效取值范围，请参考[表12](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__t27de3279a99a48968dacb015c498d9cb)。
	MetricValue string `json:"metric_value"`
	// 指标判断逻辑运算符，包括：  LT：小于 GT：大于 LTOE：小于等于 GTOE：大于等于
	ComparisonOperator *string `json:"comparison_operator,omitempty"`
	// 判断连续满足指标阈值的周期数(一个周期为5分钟)。  取值范围[1～288]
	EvaluationPeriods int32 `json:"evaluation_periods"`
}

func (o Trigger) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}
