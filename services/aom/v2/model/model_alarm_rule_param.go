package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 阈值规则实体
type AlarmRuleParam struct {
	// 是否启用通知。

	ActionEnabled *bool `json:"action_enabled,omitempty"`
	// 告警状态通知列表。

	AlarmActions *[]string `json:"alarm_actions,omitempty"`
	// 告警清除建议。

	AlarmAdvice *string `json:"alarm_advice,omitempty"`
	// 阈值规则描述。

	AlarmDescription *string `json:"alarm_description,omitempty"`
	// 告警级别。

	AlarmLevel int32 `json:"alarm_level"`
	// 阈值规则名称。

	AlarmRuleName string `json:"alarm_rule_name"`
	// 极限条件。

	ComparisonOperator string `json:"comparison_operator"`
	// 时间序列维度。

	Dimensions []Dimension `json:"dimensions"`
	// 间隔周期。

	EvaluationPeriods int32 `json:"evaluation_periods"`
	// 阈值规则是否启用。

	IdTurnOn *bool `json:"id_turn_on,omitempty"`
	// 数据不足通知列表。

	InsufficientDataActions *[]string `json:"insufficient_data_actions,omitempty"`
	// | 取值范围 名称长度为1~255个字符 时间序列名称。

	MetricName string `json:"metric_name"`
	// 时间序列命名空间。

	Namespace string `json:"namespace"`
	// 正常状态通知列表。

	OkActions *[]string `json:"ok_actions,omitempty"`
	// 统计周期。

	Period int32 `json:"period"`
	// 取值范围 \"maximum\",\"minimum\",\"average\", \"sum\",\"sampleCount\" 统计方式

	Statistic string `json:"statistic"`
	// 临界值。

	Threshold string `json:"threshold"`
	// 阈值单元。

	Unit string `json:"unit"`
}

func (o AlarmRuleParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRuleParam struct{}"
	}

	return strings.Join([]string{"AlarmRuleParam", string(data)}, " ")
}
