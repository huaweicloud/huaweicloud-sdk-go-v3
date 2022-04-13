package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询指标
type TransformMetric struct {
	// 声明属性作为表达式参数

	Inputs []InputParam `json:"inputs"`
	// 指标名称

	MetricName string `json:"metric_name"`
	// 表达式

	Expression string `json:"expression"`
}

func (o TransformMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransformMetric struct{}"
	}

	return strings.Join([]string{"TransformMetric", string(data)}, " ")
}
