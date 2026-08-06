package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportSingleValueStat struct {

	// 指标名。
	Metric *string `json:"metric,omitempty"`

	// 数值。
	Value *float64 `json:"value,omitempty"`

	// 最大值。
	MaxValue *float64 `json:"max_value,omitempty"`

	// 归一化值。
	Normalized *float64 `json:"normalized,omitempty"`

	// 当前状态。
	Stage *string `json:"stage,omitempty"`

	// 指标采集时间。
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o HealthReportSingleValueStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportSingleValueStat struct{}"
	}

	return strings.Join([]string{"HealthReportSingleValueStat", string(data)}, " ")
}
