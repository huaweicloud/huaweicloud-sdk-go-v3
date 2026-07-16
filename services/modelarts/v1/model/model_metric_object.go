package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricObject 运行指标。
type MetricObject struct {

	// 运行指标，可选值如下： - cpuUsage：CPU使用率 - memUsage：物理内存使用率 - gpuUtil：GPU使用率 - gpuMemUsage：显存使用率 - npuUtil：NPU使用率 - npuMemUsage：NPU显存使用率
	Metric *string `json:"metric,omitempty"`

	// 运行指标对应数值，1min统计一个平均值。
	Value *[]float64 `json:"value,omitempty"`
}

func (o MetricObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricObject struct{}"
	}

	return strings.Join([]string{"MetricObject", string(data)}, " ")
}
