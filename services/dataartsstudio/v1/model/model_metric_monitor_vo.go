package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricMonitorVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 其他指标ID，填写String类型替代Long类型。
	OtherMetricIds *[]string `json:"other_metric_ids,omitempty"`

	// 其他指标名称，只读。
	OtherMetricNames *[]string `json:"other_metric_names,omitempty"`

	// 其他复合指标ID。
	OtherCompoundMetricIds *[]string `json:"other_compound_metric_ids,omitempty"`

	// 其他复合指标名称。
	OtherCompoundMetricNames *[]string `json:"other_compound_metric_names,omitempty"`

	// 告警表达式。
	Expression *string `json:"expression,omitempty"`

	// 挂载指ID，填写String类型替代Long类型。
	MetricId *string `json:"metric_id,omitempty"`

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
