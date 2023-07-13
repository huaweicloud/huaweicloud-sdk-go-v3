package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregateMetrics 聚合计算定义
type AggregateMetrics struct {
	TimeSpan *TimeSpan `json:"time_span"`

	// 聚合时间间隔, 示例：\"1d|1h|10m|10s\"
	Interval *string `json:"interval,omitempty"`

	// 聚合时间偏移量, 需要小于interval, 示例： \"1h|10m|10s\"
	Offset *string `json:"offset,omitempty"`

	// 对property按指定tags标签进行过滤查询，填入设备标签与标签值，不可为空，例如 {\"deviceId\": \"id0001\"}
	Tags map[string]string `json:"tags"`

	// 查询的测量指标列表
	Metrics []AggregateMetric `json:"metrics"`

	// 返回值个数限制
	Limit *int32 `json:"limit,omitempty"`
}

func (o AggregateMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateMetrics struct{}"
	}

	return strings.Join([]string{"AggregateMetrics", string(data)}, " ")
}
