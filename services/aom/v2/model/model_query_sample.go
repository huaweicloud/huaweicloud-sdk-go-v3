package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询参数集
type QuerySample struct {
	// 时间序列的命名空间。 取值范围 PAAS.CONTAINER PAAS.NODE PAAS.SLA PAAS.AGGR CUSTOMMETRICS

	Namespace string `json:"namespace"`
	// 时间序列维度列表。 取值范围： 数组不能为空，同时数组中任何一个dimension对象name和value属性的值也不能为空。

	Dimensions []DimensionSeries `json:"dimensions"`
	// 时间序列名称。 取值范围 名称长度为1~255个字符

	MetricName string `json:"metric_name"`
}

func (o QuerySample) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySample struct{}"
	}

	return strings.Join([]string{"QuerySample", string(data)}, " ")
}
