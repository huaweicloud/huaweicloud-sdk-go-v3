package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerMetricData 监控信息。
type ServerMetricData struct {

	// 监控指标名称。
	MetricName *string `json:"metric_name,omitempty"`

	// 指标数据列表。
	Datapoints *[]ServerDataPoints `json:"datapoints,omitempty"`

	// 维度值，仅查询GPU监控信息时有值。
	DimensionValue *string `json:"dimension_value,omitempty"`
}

func (o ServerMetricData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerMetricData struct{}"
	}

	return strings.Join([]string{"ServerMetricData", string(data)}, " ")
}
