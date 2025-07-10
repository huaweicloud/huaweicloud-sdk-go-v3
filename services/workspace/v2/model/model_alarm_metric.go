package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmMetric 告警指标信息。
type AlarmMetric struct {

	// 查询服务的命名空间。
	Namespace *string `json:"namespace,omitempty"`

	// 资源的监控指标名称。
	MetricName *string `json:"metric_name,omitempty"`

	// 指标维度，目前最大可添加4个维度。
	Dimensions *[]AlarmMetricDimension `json:"dimensions,omitempty"`
}

func (o AlarmMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmMetric struct{}"
	}

	return strings.Join([]string{"AlarmMetric", string(data)}, " ")
}
