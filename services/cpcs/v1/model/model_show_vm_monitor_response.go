package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVmMonitorResponse Response Object
type ShowVmMonitorResponse struct {

	// 指标的时间序列
	Datapoints *[]Datapoint `json:"datapoints,omitempty"`

	// 指标名称，比如:cpu_util
	MetricName *string `json:"metric_name,omitempty"`

	// 最大值，未计算默认为0
	Max *float64 `json:"max,omitempty"`

	// 平均值，未计算默认为0
	Average        *float64 `json:"average,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowVmMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVmMonitorResponse struct{}"
	}

	return strings.Join([]string{"ShowVmMonitorResponse", string(data)}, " ")
}
