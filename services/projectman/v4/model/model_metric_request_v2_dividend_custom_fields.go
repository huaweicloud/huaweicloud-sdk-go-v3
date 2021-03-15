package model

import (
	"encoding/json"

	"strings"
)

type MetricRequestV2DividendCustomFields struct {
	// 自定义字段名称
	Name *string `json:"name,omitempty"`
	// 自定义字段取值，逗号分隔
	Options *string `json:"options,omitempty"`
}

func (o MetricRequestV2DividendCustomFields) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetricRequestV2DividendCustomFields struct{}"
	}

	return strings.Join([]string{"MetricRequestV2DividendCustomFields", string(data)}, " ")
}
