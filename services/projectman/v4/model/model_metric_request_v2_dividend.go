package model

import (
	"encoding/json"

	"strings"
)

// 指标分子过滤条件
type MetricRequestV2Dividend struct {
	// 自定义类型过滤条件

	CustomFields *[]MetricRequestV2DividendCustomFields `json:"custom_fields,omitempty"`
}

func (o MetricRequestV2Dividend) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetricRequestV2Dividend struct{}"
	}

	return strings.Join([]string{"MetricRequestV2Dividend", string(data)}, " ")
}
