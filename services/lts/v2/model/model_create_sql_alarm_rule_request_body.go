package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSqlAlarmRuleRequestBody struct {

	// SQL告警名称
	SqlAlarmRuleName string `json:"sql_alarm_rule_name"`

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
	SqlAlarmLevel CreateSqlAlarmRuleRequestBodySqlAlarmLevel `json:"sql_alarm_level"`

	// 是否发送
	SqlAlarmSend bool `json:"sql_alarm_send"`

	// domainId
	DomainId string `json:"domain_id"`

	NotificationSaveRule *SqlNotificationSaveRule `json:"notification_save_rule,omitempty"`

	// 触发条件：触发次数;默认为1
	TriggerConditionCount *int32 `json:"trigger_condition_count,omitempty"`

	// 触发条件：触发周期;默认为1
	TriggerConditionFrequency *int32 `json:"trigger_condition_frequency,omitempty"`

	// 是否打开恢复通知;默认false
	WhetherRecoveryPolicy *bool `json:"whether_recovery_policy,omitempty"`

	// 恢复策略周期;默认为3
	RecoveryPolicy *int32 `json:"recovery_policy,omitempty"`

	// 通知频率,单位(分钟)
	NotificationFrequency CreateSqlAlarmRuleRequestBodyNotificationFrequency `json:"notification_frequency"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`
}

func (o CreateSqlAlarmRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlAlarmRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlAlarmRuleRequestBody", string(data)}, " ")
}

type CreateSqlAlarmRuleRequestBodySqlAlarmLevel struct {
	value string
}

type CreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum struct {
	INFO     CreateSqlAlarmRuleRequestBodySqlAlarmLevel
	MINOR    CreateSqlAlarmRuleRequestBodySqlAlarmLevel
	MAJOR    CreateSqlAlarmRuleRequestBodySqlAlarmLevel
	CRITICAL CreateSqlAlarmRuleRequestBodySqlAlarmLevel
}

func GetCreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum() CreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum {
	return CreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum{
		INFO: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Info",
		},
		MINOR: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Minor",
		},
		MAJOR: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Major",
		},
		CRITICAL: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Critical",
		},
	}
}

func (c CreateSqlAlarmRuleRequestBodySqlAlarmLevel) Value() string {
	return c.value
}

func (c CreateSqlAlarmRuleRequestBodySqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlAlarmRuleRequestBodySqlAlarmLevel) UnmarshalJSON(b []byte) error {
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

type CreateSqlAlarmRuleRequestBodyNotificationFrequency struct {
	value int32
}

type CreateSqlAlarmRuleRequestBodyNotificationFrequencyEnum struct {
	E_0   CreateSqlAlarmRuleRequestBodyNotificationFrequency
	E_5   CreateSqlAlarmRuleRequestBodyNotificationFrequency
	E_10  CreateSqlAlarmRuleRequestBodyNotificationFrequency
	E_15  CreateSqlAlarmRuleRequestBodyNotificationFrequency
	E_30  CreateSqlAlarmRuleRequestBodyNotificationFrequency
	E_60  CreateSqlAlarmRuleRequestBodyNotificationFrequency
	E_180 CreateSqlAlarmRuleRequestBodyNotificationFrequency
	E_360 CreateSqlAlarmRuleRequestBodyNotificationFrequency
}

func GetCreateSqlAlarmRuleRequestBodyNotificationFrequencyEnum() CreateSqlAlarmRuleRequestBodyNotificationFrequencyEnum {
	return CreateSqlAlarmRuleRequestBodyNotificationFrequencyEnum{
		E_0: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 0,
		}, E_5: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 5,
		}, E_10: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 10,
		}, E_15: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 15,
		}, E_30: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 30,
		}, E_60: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 60,
		}, E_180: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 180,
		}, E_360: CreateSqlAlarmRuleRequestBodyNotificationFrequency{
			value: 360,
		},
	}
}

func (c CreateSqlAlarmRuleRequestBodyNotificationFrequency) Value() int32 {
	return c.value
}

func (c CreateSqlAlarmRuleRequestBodyNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlAlarmRuleRequestBodyNotificationFrequency) UnmarshalJSON(b []byte) error {
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
