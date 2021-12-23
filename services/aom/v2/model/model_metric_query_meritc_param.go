package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询参数集
type MetricQueryMeritcParam struct {
	// 指标维度列表。 取值范围： 数组不能为空，同时数组中任何一个dimension对象name和value属性的值也不能为空。

	Dimensions []Dimension `json:"dimensions"`
	// 指标名称。 取值范围 名称长度为1~255个字符

	MetricName string `json:"metricName"`
	// 指标命名空间。 取值范围 PAAS.CONTAINER PAAS.NODE PAAS.SLA PAAS.AGGR CUSTOMMETRICS

	Namespace string `json:"namespace"`
}

func (o MetricQueryMeritcParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricQueryMeritcParam struct{}"
	}

	return strings.Join([]string{"MetricQueryMeritcParam", string(data)}, " ")
}
