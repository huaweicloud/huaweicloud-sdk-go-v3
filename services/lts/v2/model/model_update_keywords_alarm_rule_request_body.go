package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateKeywordsAlarmRuleRequestBody struct {

	// 关键词告警规则id
	KeywordsAlarmRuleId string `json:"keywords_alarm_rule_id"`

	// 关键词告警名称
	KeywordsAlarmRuleName string `json:"keywords_alarm_rule_name"`

	// 规则名称
	AlarmRuleAlias *string `json:"alarm_rule_alias,omitempty"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription *string `json:"keywords_alarm_rule_description,omitempty"`

	// 关键词详细信息
	KeywordsRequests []KeywordsRequest `json:"keywords_requests"`

	Frequency *Frequency `json:"frequency"`

	// 告警级别
	KeywordsAlarmLevel UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel `json:"keywords_alarm_level"`

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
	NotificationFrequency UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency `json:"notification_frequency"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`

	// **参数解释：** 告警标签信息。标签是以键值对（key-value）的形式表示，key和value为一一对应关系。 **约束限制：** 不涉及。
	Tags *[]TagsRequestBody `json:"tags,omitempty"`

	// **参数解释：** 企业项目id信息。 **约束限制：** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o UpdateKeywordsAlarmRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeywordsAlarmRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeywordsAlarmRuleRequestBody", string(data)}, " ")
}

type UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel struct {
	value string
}

type UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum struct {
	INFO     UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
	MINOR    UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
	MAJOR    UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
	CRITICAL UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel
}

func GetUpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum() UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum {
	return UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevelEnum{
		INFO: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Info",
		},
		MINOR: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Minor",
		},
		MAJOR: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Major",
		},
		CRITICAL: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel{
			value: "Critical",
		},
	}
}

func (c UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel) Value() string {
	return c.value
}

func (c UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel) UnmarshalJSON(b []byte) error {
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

type UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency struct {
	value int32
}

type UpdateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum struct {
	E_0   UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_5   UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_10  UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_15  UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_30  UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_60  UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_180 UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
	E_360 UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency
}

func GetUpdateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum() UpdateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum {
	return UpdateKeywordsAlarmRuleRequestBodyNotificationFrequencyEnum{
		E_0: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 0,
		}, E_5: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 5,
		}, E_10: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 10,
		}, E_15: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 15,
		}, E_30: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 30,
		}, E_60: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 60,
		}, E_180: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 180,
		}, E_360: UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency{
			value: 360,
		},
	}
}

func (c UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency) Value() int32 {
	return c.value
}

func (c UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKeywordsAlarmRuleRequestBodyNotificationFrequency) UnmarshalJSON(b []byte) error {
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
