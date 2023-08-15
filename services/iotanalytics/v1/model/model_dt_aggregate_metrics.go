package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DtAggregateMetrics 聚合指标列表
type DtAggregateMetrics struct {

	// 定义指标计算查询的输入资产属性列表
	Inputs []MetricInput `json:"inputs"`

	// 指标表达式,最多64个字符
	Expression string `json:"expression"`

	// 指标名，指标计算查询的输出指标名称
	MetricName string `json:"metric_name"`
}

func (o DtAggregateMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DtAggregateMetrics struct{}"
	}

	return strings.Join([]string{"DtAggregateMetrics", string(data)}, " ")
}
