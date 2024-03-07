package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddMetricNotifyRuleReq 设置通知规则
type AddMetricNotifyRuleReq struct {

	// 统计指标名称，目前仅支持固定值：desktop_idle_duration 同一指标的规则不允许重复 * `desktop_idle_duration` -  桌面空闲时长, 仅允许设置 '>=' threshold
	MetricName string `json:"metric_name"`

	// 规则配置-阈值(天)
	Threshold *int32 `json:"threshold,omitempty"`

	// 统计指标对应的统计值和threshold进行比较的条件 * `>=` -  统计指标大于等于threshold时触发 * `>` -   统计指标大于threshold时触发 * `=` -  统计指标等于threshold时触发 * `<=` -  统计指标小于等于threshold时触发 * `<` -  统计指标小于threshold时触发
	ComparisonOperator string `json:"comparison_operator"`

	// 触发通知后；下次通知的间隔时间;默认每天一次
	Interval *int32 `json:"interval,omitempty"`

	// 启禁用规则 true:启用 false:禁用
	Enable bool `json:"enable"`

	// 通知对象;smn的主题urn
	NotifyObject string `json:"notify_object"`
}

func (o AddMetricNotifyRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMetricNotifyRuleReq struct{}"
	}

	return strings.Join([]string{"AddMetricNotifyRuleReq", string(data)}, " ")
}
