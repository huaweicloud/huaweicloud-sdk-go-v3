package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数项。
type QuerySeriesOptionParam struct {
	// 取值范围 PAAS.CONTAINER、PAAS.NODE、PAAS.SLA、PAAS.AGGR、CUSTOMMETRICS等 时间序列命名空间。 PAAS.CONTAINER：应用时间序列； PAAS.NODE：节时间序列； PAAS.SLA：SLA时间序列； PAAS.AGGR：集群时间序列； CUSTOMMETRICS：自定义时间序列

	Namespace string `json:"namespace"`
	//  | 取值范围 名称长度为1~255个字符 时间序列名称。

	MetricName *string `json:"metric_name,omitempty"`
	// 时间序列维度列表。

	Dimensions *[]DimensionSeries `json:"dimensions,omitempty"`
}

func (o QuerySeriesOptionParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySeriesOptionParam struct{}"
	}

	return strings.Join([]string{"QuerySeriesOptionParam", string(data)}, " ")
}
