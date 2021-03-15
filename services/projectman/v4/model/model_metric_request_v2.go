package model

import (
	"encoding/json"

	"strings"
)

type MetricRequestV2 struct {
	// 统计周期
	DateRange *string `json:"date_range,omitempty"`
	// 指标类型
	MetricType *string                  `json:"metric_type,omitempty"`
	Dividend   *MetricRequestV2Dividend `json:"dividend,omitempty"`
	// 指标分母过滤条件
	Divisor *interface{} `json:"divisor,omitempty"`
}

func (o MetricRequestV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetricRequestV2 struct{}"
	}

	return strings.Join([]string{"MetricRequestV2", string(data)}, " ")
}
