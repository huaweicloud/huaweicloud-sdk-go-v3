package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricExtraInfo 告警策略附加信息，默认为空。
type MetricExtraInfo struct {

	// 原始指标名称
	OriginMetricName string `json:"origin_metric_name"`

	// 指标名称前缀
	MetricPrefix *string `json:"metric_prefix,omitempty"`

	// 用户进程名称
	CustomProcName *string `json:"custom_proc_name,omitempty"`

	// 指标类型
	MetricType *string `json:"metric_type,omitempty"`
}

func (o MetricExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricExtraInfo struct{}"
	}

	return strings.Join([]string{"MetricExtraInfo", string(data)}, " ")
}
