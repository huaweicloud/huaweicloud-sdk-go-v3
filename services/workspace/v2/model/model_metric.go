package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metric 指标
type Metric struct {

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 指标值
	MetricValue *float64 `json:"metric_value,omitempty"`
}

func (o Metric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metric struct{}"
	}

	return strings.Join([]string{"Metric", string(data)}, " ")
}
