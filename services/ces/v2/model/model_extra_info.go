package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtraInfo struct {

	// 指标名称
	OriginMetricName *string `json:"origin_metric_name,omitempty"`

	// 指标名称前缀
	MetricPrefix *string `json:"metric_prefix,omitempty"`

	// 指标类型
	MetricType *string `json:"metric_type,omitempty"`

	// 自定义进程名称
	CustomProcName *string `json:"custom_proc_name,omitempty"`
}

func (o ExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraInfo struct{}"
	}

	return strings.Join([]string{"ExtraInfo", string(data)}, " ")
}
