package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metric 作业的所有输入输出监控信息。
type Metric struct {

	// 所有输入流。
	Sources *[]FlinkJobMetricInfo `json:"sources,omitempty"`

	// 所有输出流。
	Sinks *[]FlinkJobMetricInfo `json:"sinks,omitempty"`

	// 总输入速率
	TotalReadRate *float64 `json:"total_read_rate,omitempty"`

	// 总输出速率
	TotalWriteRate *float64 `json:"total_write_rate,omitempty"`
}

func (o Metric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metric struct{}"
	}

	return strings.Join([]string{"Metric", string(data)}, " ")
}
