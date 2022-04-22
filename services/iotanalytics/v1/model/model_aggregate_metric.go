package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询指标
type AggregateMetric struct {

	// 声明属性作为表达式参数
	Inputs []InputParam `json:"inputs"`

	// 指标名称
	MetricName string `json:"metric_name"`

	// 表达式
	Expression string `json:"expression"`
}

func (o AggregateMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateMetric struct{}"
	}

	return strings.Join([]string{"AggregateMetric", string(data)}, " ")
}
