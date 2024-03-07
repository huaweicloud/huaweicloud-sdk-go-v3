package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricNotifyRule 指标通知规则
type MetricNotifyRule struct {

	// 规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 统计指标名称，目前仅支持固定值：desktop_idle_duration * `desktop_idle_duration` -  桌面空闲时长
	MetricName *string `json:"metric_name,omitempty"`

	// 统计持续周期(天)
	Threshold *int32 `json:"threshold,omitempty"`

	// 统计指标对应的统计值和threshold进行比较的条件 * `>=` -  统计指标大于等于threshold时触发 * `>` -   统计指标大于threshold时触发 * `=` -  统计指标等于threshold时触发 * `<=` -  统计指标小于等于threshold时触发 * `<` -  统计指标小于threshold时触发
	ComparisonOperator *string `json:"comparison_operator,omitempty"`

	// 触发通知后；下次通知的间隔时间;默认每天一次
	Interval *int32 `json:"interval,omitempty"`

	// 启禁用规则 true:启用 false:禁用
	Enable *bool `json:"enable,omitempty"`

	// 通知对象;smn的主题urn
	NotifyObject *string `json:"notify_object,omitempty"`
}

func (o MetricNotifyRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricNotifyRule struct{}"
	}

	return strings.Join([]string{"MetricNotifyRule", string(data)}, " ")
}
