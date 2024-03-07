package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMetricNotifyRuleReq 设置通知规则
type UpdateMetricNotifyRuleReq struct {

	// 统计指标名称，目前仅支持固定值：desktop_idle_duration 同一指标的规则不允许重复 * `desktop_idle_duration` -  桌面空闲时长
	MetricName *string `json:"metric_name,omitempty"`

	// 规则配置-阈值(天)
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

func (o UpdateMetricNotifyRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMetricNotifyRuleReq struct{}"
	}

	return strings.Join([]string{"UpdateMetricNotifyRuleReq", string(data)}, " ")
}
