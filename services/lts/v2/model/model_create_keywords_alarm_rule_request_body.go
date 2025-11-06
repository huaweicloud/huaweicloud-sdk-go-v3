package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateKeywordsAlarmRuleRequestBody struct {

	// 关键词告警名称  >不能以点和下划线开头或以点结尾。
	KeywordsAlarmRuleName string `json:"keywords_alarm_rule_name"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription *string `json:"keywords_alarm_rule_description,omitempty"`

	// 关键词详细信息
	KeywordsRequests []KeywordsRequest `json:"keywords_requests"`

	Frequency *Frequency `json:"frequency"`

	// 告警级别
	KeywordsAlarmLevel CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel `json:"keywords_alarm_level"`

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
	NotificationFrequency CreateKeywordsAlarmRuleRequestBodyNotificationFrequency `json:"notification_frequency"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`

	// **参数解释：** 告警标签信息。标签是以键值对（key-value）的形式表示，key和value为一一对应关系。 **约束限制：** 不涉及。
	Tags *[]TagsRequestBody `json:"tags,omitempty"`
}

func (o CreateKeywordsAlarmRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeywordsAlarmRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeywordsAlarmRuleRequestBody", string(data)}, " ")
}

type CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel struct {
	value string
}

type CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum struct {
	INFO     CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
	MINOR    CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
	MAJOR    CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
	CRITICAL CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
}

func GetCreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum() CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum {
	return CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum{
		INFO: CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Info",
		},
		MINOR: CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Minor",
		},
		MAJOR: CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Major",
		},
		CRITICAL: CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Critical",
		},
	}
}

func (c CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel) Value() string {
	return c.value
}

func (c CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel) UnmarshalJSON(b []byte) error {
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

type CreateKeywordsAlarmRuleRequestBodyNotificationFrequency struct {
	value int32
}

type CreateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum struct {
	E_0   CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_5   CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_10  CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_15  CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_30  CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_60  CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_180 CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_360 CreateKeywordsAlarmRuleRequestBodyNotificationFrequency
}

func GetCreateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum() CreateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum {
	return CreateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum{
		E_0: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 0,
		}, E_5: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 5,
		}, E_10: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 10,
		}, E_15: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 15,
		}, E_30: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 30,
		}, E_60: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 60,
		}, E_180: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 180,
		}, E_360: CreateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 360,
		},
	}
}

func (c CreateKeywordsAlarmRuleRequestBodyNotificationFrequency) Value() int32 {
	return c.value
}

func (c CreateKeywordsAlarmRuleRequestBodyNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeywordsAlarmRuleRequestBodyNotificationFrequency) UnmarshalJSON(b []byte) error {
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
