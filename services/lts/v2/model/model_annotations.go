package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Annotations struct {

	// 告警列表详情
	Message string `json:"message"`

	// 日志组/流id,名称
	LogInfo string `json:"log_info"`

	// 当前值
	CurrentValue string `json:"current_value"`

	// (sql/关键词)告警详情原始数据
	OldAnnotations string `json:"old_annotations"`

	// **参数解释：** 告警行动规则名称。 **取值范围：** 不涉及。
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`

	// **参数解释：** 告警规则别名。 **取值范围：** 不涉及。
	AlarmRuleAlias *string `json:"alarm_rule_alias,omitempty"`

	// **参数解释：** 告警规则url。 **取值范围：** 不涉及。
	AlarmRuleUrl *string `json:"alarm_rule_url,omitempty"`

	// **参数解释：** 告警触发状态。 **取值范围：** - 触发(firing) - 恢复(resolved)
	AlarmStatus *string `json:"alarm_status,omitempty"`

	// **参数解释：** 告警触发条件表达式。 **取值范围：** 不涉及。
	ConditionExpression *string `json:"condition_expression,omitempty"`

	// **参数解释：** 告警触发表达式带值。例如：条件表达式为pv > 0，则condition_expression_with_value取值为：100 > 0。 **取值范围：** 不涉及。
	ConditionExpressionWithValue *string `json:"condition_expression_with_value,omitempty"`

	// **参数解释：** 通知频率。 **取值范围：** 不涉及。
	NotificationFrequency *string `json:"notification_frequency,omitempty"`

	// **参数解释：** 告警ID。 **取值范围：** 不涉及。
	RecordId *string `json:"record_id,omitempty"`

	// **参数解释：** 是否开启告警恢复开关。 **取值范围：** - true： 开启告警恢复 - false： 关闭告警恢复
	RecoveryPolicy *bool `json:"recovery_policy,omitempty"`

	// **参数解释：** 告警通知的详细信息。
	Results *[]Results `json:"results,omitempty"`

	// **参数解释：** 告警统计周期。 **取值范围：** 不涉及。
	Frequency *string `json:"frequency,omitempty"`

	// **参数解释：** 告警规则类型。 **取值范围：** - sql： sql告警 - keywords： 关键词告警
	Type *string `json:"type,omitempty"`
}

func (o Annotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Annotations struct{}"
	}

	return strings.Join([]string{"Annotations", string(data)}, " ")
}
