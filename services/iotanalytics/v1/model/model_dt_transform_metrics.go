package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转换指标列表
type DtTransformMetrics struct {
	// 定义指标计算查询的输入资产属性列表

	Inputs []MetricInput `json:"inputs"`
	// 指标表达式,最多64个字符

	Expression string `json:"expression"`
	// 指标名，指标计算查询的输出指标名称

	MetricName string `json:"metric_name"`
}

func (o DtTransformMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DtTransformMetrics struct{}"
	}

	return strings.Join([]string{"DtTransformMetrics", string(data)}, " ")
}
