package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSqlAlarmRuleResponse Response Object
type UpdateSqlAlarmRuleResponse struct {

	// SQL告警名称
	SqlAlarmRuleName *string `json:"sql_alarm_rule_name,omitempty"`

	// 规则名称
	AlarmRuleAlias *string `json:"alarm_rule_alias,omitempty"`

	// 是否管道符sql查询
	IsCssSql *bool `json:"is_css_sql,omitempty"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty"`

	// SQL告警规则id
	SqlAlarmRuleId *string `json:"sql_alarm_rule_id,omitempty"`

	// SQL告警信息描述
	SqlAlarmRuleDescription *string `json:"sql_alarm_rule_description,omitempty"`

	// SQL详细信息
	SqlRequests *[]SqlRequestResponse `json:"sql_requests,omitempty"`

	Frequency *FrequencyRespBody `json:"frequency,omitempty"`

	// 条件表达式
	ConditionExpression *string `json:"condition_expression,omitempty"`

	// 告警级别
	SqlAlarmLevel *UpdateSqlAlarmRuleResponseSqlAlarmLevel `json:"sql_alarm_level,omitempty"`

	// 是否发送
	SqlAlarmSend *bool `json:"sql_alarm_send,omitempty"`

	// domainId
	DomainId *string `json:"domain_id,omitempty"`

	// 通知频率,单位(分钟)
	NotificationFrequency *UpdateSqlAlarmRuleResponseNotificationFrequency `json:"notification_frequency,omitempty"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`

	// **参数解释：** 告警触发条件：满足条件次数。满足条件次数是指设置的SQL语句。 **取值范围：** 不涉及。
	TriggerConditionCount *int32 `json:"trigger_condition_count,omitempty"`

	// **参数解释：** 触发条件：触发周期。 **取值范围：** 不涉及。
	TriggerConditionFrequency *int32 `json:"trigger_condition_frequency,omitempty"`

	// **参数解释：** 是否打开恢复通知。 **取值范围：** - true：配置告警恢复策略。 - false：不配置告警恢复策略。
	WhetherRecoveryPolicy *bool `json:"whether_recovery_policy,omitempty"`

	// **参数解释：** 恢复策略周期。 **取值范围：** 不涉及。
	RecoveryPolicy *int32 `json:"recovery_policy,omitempty"`

	// **参数解释：** 告警标签信息。标签是以键值对（key-value）的形式表示，key和value为一一对应关系。
	Tags           *[]TagsResBody `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateSqlAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateSqlAlarmRuleResponse", string(data)}, " ")
}

type UpdateSqlAlarmRuleResponseSqlAlarmLevel struct {
	value string
}

type UpdateSqlAlarmRuleResponseSqlAlarmLevelEnum struct {
	INFO     UpdateSqlAlarmRuleResponseSqlAlarmLevel
	MINOR    UpdateSqlAlarmRuleResponseSqlAlarmLevel
	MAJOR    UpdateSqlAlarmRuleResponseSqlAlarmLevel
	CRITICAL UpdateSqlAlarmRuleResponseSqlAlarmLevel
}

func GetUpdateSqlAlarmRuleResponseSqlAlarmLevelEnum() UpdateSqlAlarmRuleResponseSqlAlarmLevelEnum {
	return UpdateSqlAlarmRuleResponseSqlAlarmLevelEnum{
		INFO: UpdateSqlAlarmRuleResponseSqlAlarmLevel{
			value: "Info",
		},
		MINOR: UpdateSqlAlarmRuleResponseSqlAlarmLevel{
			value: "Minor",
		},
		MAJOR: UpdateSqlAlarmRuleResponseSqlAlarmLevel{
			value: "Major",
		},
		CRITICAL: UpdateSqlAlarmRuleResponseSqlAlarmLevel{
			value: "CRITICAL",
		},
	}
}

func (c UpdateSqlAlarmRuleResponseSqlAlarmLevel) Value() string {
	return c.value
}

func (c UpdateSqlAlarmRuleResponseSqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlAlarmRuleResponseSqlAlarmLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateSqlAlarmRuleResponseNotificationFrequency struct {
	value int32
}

type UpdateSqlAlarmRuleResponseNotificationFrequencyEnum struct {
	E_0   UpdateSqlAlarmRuleResponseNotificationFrequency
	E_5   UpdateSqlAlarmRuleResponseNotificationFrequency
	E_10  UpdateSqlAlarmRuleResponseNotificationFrequency
	E_15  UpdateSqlAlarmRuleResponseNotificationFrequency
	E_30  UpdateSqlAlarmRuleResponseNotificationFrequency
	E_60  UpdateSqlAlarmRuleResponseNotificationFrequency
	E_180 UpdateSqlAlarmRuleResponseNotificationFrequency
	E_360 UpdateSqlAlarmRuleResponseNotificationFrequency
}

func GetUpdateSqlAlarmRuleResponseNotificationFrequencyEnum() UpdateSqlAlarmRuleResponseNotificationFrequencyEnum {
	return UpdateSqlAlarmRuleResponseNotificationFrequencyEnum{
		E_0: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 0,
		}, E_5: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 5,
		}, E_10: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 10,
		}, E_15: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 15,
		}, E_30: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 30,
		}, E_60: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 60,
		}, E_180: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 180,
		}, E_360: UpdateSqlAlarmRuleResponseNotificationFrequency{
			value: 360,
		},
	}
}

func (c UpdateSqlAlarmRuleResponseNotificationFrequency) Value() int32 {
	return c.value
}

func (c UpdateSqlAlarmRuleResponseNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlAlarmRuleResponseNotificationFrequency) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
