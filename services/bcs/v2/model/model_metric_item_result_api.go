package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监控数据信息
type MetricItemResultApi struct {
	Metric *MetricDemision `json:"metric,omitempty" xml:"metric"`

	// 监控数据信息
	DataPoints *[]MetricDataPoints `json:"dataPoints,omitempty" xml:"dataPoints"`
}

func (o MetricItemResultApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricItemResultApi struct{}"
	}

	return strings.Join([]string{"MetricItemResultApi", string(data)}, " ")
}
