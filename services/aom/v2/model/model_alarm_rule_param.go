package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 阈值规则实体
type AlarmRuleParam struct {

	// 是否启用通知。
	ActionEnabled *bool `json:"action_enabled,omitempty" xml:"action_enabled"`

	// 告警状态通知列表。
	AlarmActions *[]string `json:"alarm_actions,omitempty" xml:"alarm_actions"`

	// 告警清除建议。
	AlarmAdvice *string `json:"alarm_advice,omitempty" xml:"alarm_advice"`

	// 阈值规则描述。
	AlarmDescription *string `json:"alarm_description,omitempty" xml:"alarm_description"`

	// 告警级别。1：紧急，2：重要，3：一般，4：提示。
	AlarmLevel int32 `json:"alarm_level" xml:"alarm_level"`

	// 阈值规则名称。
	AlarmRuleName string `json:"alarm_rule_name" xml:"alarm_rule_name"`

	// 超限条件。
	ComparisonOperator string `json:"comparison_operator" xml:"comparison_operator"`

	// 时间序列维度。
	Dimensions []Dimension `json:"dimensions" xml:"dimensions"`

	// 间隔周期。
	EvaluationPeriods int32 `json:"evaluation_periods" xml:"evaluation_periods"`

	// 阈值规则是否启用。
	IdTurnOn *bool `json:"id_turn_on,omitempty" xml:"id_turn_on"`

	// 数据不足通知列表。
	InsufficientDataActions *[]string `json:"insufficient_data_actions,omitempty" xml:"insufficient_data_actions"`

	// 时间序列名称。名称长度取值范围为1~255个字符。
	MetricName string `json:"metric_name" xml:"metric_name"`

	// 时间序列命名空间。
	Namespace string `json:"namespace" xml:"namespace"`

	// 正常状态通知列表。
	OkActions *[]string `json:"ok_actions,omitempty" xml:"ok_actions"`

	// 统计周期。
	Period int32 `json:"period" xml:"period"`

	// 统计方式。
	Statistic string `json:"statistic" xml:"statistic"`

	// 超限值。
	Threshold string `json:"threshold" xml:"threshold"`

	// 时间序列单位
	Unit string `json:"unit" xml:"unit"`
}

func (o AlarmRuleParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRuleParam struct{}"
	}

	return strings.Join([]string{"AlarmRuleParam", string(data)}, " ")
}
