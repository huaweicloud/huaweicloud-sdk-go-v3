package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricMonitorVo struct {

	// 编码。
	Id *int64 `json:"id,omitempty"`

	// 其他指标ID。
	OtherMetricIds *[]int64 `json:"other_metric_ids,omitempty"`

	// 其他指标名称。
	OtherMetricNames *[]string `json:"other_metric_names,omitempty"`

	// 告警表达式。
	Expression *string `json:"expression,omitempty"`

	// 挂载指ID。
	MetricId *int64 `json:"metric_id,omitempty"`

	// 前端表达式配置，用于前端数据恢复。
	FrontConfigs *string `json:"front_configs,omitempty"`

	MetricType *BizTypeEnum `json:"metric_type,omitempty"`
}

func (o MetricMonitorVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricMonitorVo struct{}"
	}

	return strings.Join([]string{"MetricMonitorVo", string(data)}, " ")
}
