package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricThresholdItem MetricThresholdItem对象
type MetricThresholdItem struct {

	// 指标码
	MetricCode *string `json:"metric_code,omitempty"`

	// 指标名
	MetricName *string `json:"metric_name,omitempty"`

	// 阈值
	Threshold *float64 `json:"threshold,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`
}

func (o MetricThresholdItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricThresholdItem struct{}"
	}

	return strings.Join([]string{"MetricThresholdItem", string(data)}, " ")
}
