package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询结果详细。
type MetricDataValue struct {

	// 重点指标。
	DataPoints *[]MetricDataPoints `json:"dataPoints,omitempty" xml:"dataPoints"`

	Metric *MetricQueryMeritcParam `json:"metric,omitempty" xml:"metric"`
}

func (o MetricDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataValue struct{}"
	}

	return strings.Join([]string{"MetricDataValue", string(data)}, " ")
}
