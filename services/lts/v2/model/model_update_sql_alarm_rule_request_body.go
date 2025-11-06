package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateSqlAlarmRuleRequestBody struct {

	// SQL告警id
	SqlAlarmRuleId string `json:"sql_alarm_rule_id"`

	// SQL告警名称
	SqlAlarmRuleName string `json:"sql_alarm_rule_name"`

	// 规则名称
	AlarmRuleAlias *string `json:"alarm_rule_alias,omitempty"`

	// 是否管道符sql查询
	IsCssSql *bool `json:"is_css_sql,omitempty"`

	// SQL告警信息描述
	SqlAlarmRuleDescription *string `json:"sql_alarm_rule_description,omitempty"`

	// SQL详细信息
	SqlRequests []SqlRequest `json:"sql_requests"`

	Frequency *CreateSqlAlarmRuleFrequency `json:"frequency"`

	// 条件表达式
	ConditionExpression string `json:"condition_expression"`

	// 告警级别
	SqlAlarmLevel UpdateSqlAlarmRuleRequestBodySqlAlarmLevel `json:"sql_alarm_level"`

	// domainId
	DomainId string `json:"domain_id"`

	// 触发条件：触发次数;默认为1
	TriggerConditionCount *int32 `json:"trigger_condition_count,omitempty"`

	// 触发条件：触发周期;默认为1
	TriggerConditionFrequency *int32 `json:"trigger_condition_frequency,omitempty"`

	// 是否打开恢复通知;默认false
	WhetherRecoveryPolicy *bool `json:"whether_recovery_policy,omitempty"`

	// 恢复策略周期;默认为3
	RecoveryPolicy *int32 `json:"recovery_policy,omitempty"`

	// 通知频率,单位(分钟)
	NotificationFrequency UpdateSqlAlarmRuleRequestBodyNotificationFrequency `json:"notification_frequency"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`

	// **参数解释：** 告警标签信息。标签是以键值对（key-value）的形式表示，key和value为一一对应关系。 **约束限制：** 不涉及。
	Tags *[]TagsRequestBody `json:"tags,omitempty"`
}

func (o UpdateSqlAlarmRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlAlarmRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSqlAlarmRuleRequestBody", string(data)}, " ")
}

type UpdateSqlAlarmRuleRequestBodySqlAlarmLevel struct {
	value string
}

type UpdateSqlAlarmRuleRequestBodySqlAlarmLevelEnum struct {
	INFO     UpdateSqlAlarmRuleRequestBodySqlAlarmLevel
	MINOR    UpdateSqlAlarmRuleRequestBodySqlAlarmLevel
	MAJOR    UpdateSqlAlarmRuleRequestBodySqlAlarmLevel
	CRITICAL UpdateSqlAlarmRuleRequestBodySqlAlarmLevel
}

func GetUpdateSqlAlarmRuleRequestBodySqlAlarmLevelEnum() UpdateSqlAlarmRuleRequestBodySqlAlarmLevelEnum {
	return UpdateSqlAlarmRuleRequestBodySqlAlarmLevelEnum{
		INFO: UpdateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Info",
		},
		MINOR: UpdateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Minor",
		},
		MAJOR: UpdateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Major",
		},
		CRITICAL: UpdateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Critical",
		},
	}
}

func (c UpdateSqlAlarmRuleRequestBodySqlAlarmLevel) Value() string {
	return c.value
}

func (c UpdateSqlAlarmRuleRequestBodySqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlAlarmRuleRequestBodySqlAlarmLevel) UnmarshalJSON(b []byte) error {
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

type UpdateSqlAlarmRuleRequestBodyNotificationFrequency struct {
	value int32
}

type UpdateSqlAlarmRuleRequestBodyNotificationFrequencyEnum struct {
	E_0   UpdateSqlAlarmRuleRequestBodyNotificationFrequency
	E_5   UpdateSqlAlarmRuleRequestBodyNotificationFrequency
	E_10  UpdateSqlAlarmRuleRequestBodyNotificationFrequency
	E_15  UpdateSqlAlarmRuleRequestBodyNotificationFrequency
	E_30  UpdateSqlAlarmRuleRequestBodyNotificationFrequency
	E_60  UpdateSqlAlarmRuleRequestBodyNotificationFrequency
	E_180 UpdateSqlAlarmRuleRequestBodyNotificationFrequency
	E_360 UpdateSqlAlarmRuleRequestBodyNotificationFrequency
}

func GetUpdateSqlAlarmRuleRequestBodyNotificationFrequencyEnum() UpdateSqlAlarmRuleRequestBodyNotificationFrequencyEnum {
	return UpdateSqlAlarmRuleRequestBodyNotificationFrequencyEnum{
		E_0: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 0,
		}, E_5: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 5,
		}, E_10: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 10,
		}, E_15: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 15,
		}, E_30: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 30,
		}, E_60: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 60,
		}, E_180: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 180,
		}, E_360: UpdateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 360,
		},
	}
}

func (c UpdateSqlAlarmRuleRequestBodyNotificationFrequency) Value() int32 {
	return c.value
}

func (c UpdateSqlAlarmRuleRequestBodyNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlAlarmRuleRequestBodyNotificationFrequency) UnmarshalJSON(b []byte) error {
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
