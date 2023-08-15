package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateKeywordsAlarmRuleRequestBody struct {

	// 关键词告警规则id
	KeywordsAlarmRuleId string `json:"keywords_alarm_rule_id"`

	// 关键词告警名称
	KeywordsAlarmRuleName string `json:"keywords_alarm_rule_name"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription *string `json:"keywords_alarm_rule_description,omitempty"`

	// 关键词详细信息
	KeywordsRequests []KeywordsRequest `json:"keywords_requests"`

	// 告警统计周期
	Frequency *Frequency `json:"frequency"`

	// 告警级别
	KeywordsAlarmLevel UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel `json:"keywords_alarm_level"`

	// 是否发送
	KeywordsAlarmSend bool `json:"keywords_alarm_send"`

	// 发送主题 0:不变 1:新增 2:修改 3:删除
	KeywordsAlarmSendCode UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode `json:"keywords_alarm_send_code"`

	// domainId
	DomainId string `json:"domain_id"`

	// 通知主题
	NotificationSaveRule *SqlNotificationSaveRule `json:"notification_save_rule,omitempty"`

	// 触发条件：触发次数;默认为1
	TriggerConditionCount *int32 `json:"trigger_condition_count,omitempty"`

	// 触发条件：触发周期;默认为1
	TriggerConditionFrequency *int32 `json:"trigger_condition_frequency,omitempty"`

	// 是否打开恢复通知;默认false
	WhetherRecoveryPolicy *bool `json:"whether_recovery_policy,omitempty"`

	// 恢复策略周期;默认为3
	RecoveryPolicy *int32 `json:"recovery_policy,omitempty"`
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

type UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode struct {
	value int32
}

type UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCodeEnum struct {
	E_0 UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode
	E_1 UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode
	E_2 UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode
	E_3 UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode
}

func GetUpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCodeEnum() UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCodeEnum {
	return UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCodeEnum{
		E_0: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode{
			value: 0,
		}, E_1: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode{
			value: 1,
		}, E_2: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode{
			value: 2,
		}, E_3: UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode{
			value: 3,
		},
	}
}

func (c UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode) Value() int32 {
	return c.value
}

func (c UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKeywordsAlarmRuleRequestBodyKeywordsAlarmSendCode) UnmarshalJSON(b []byte) error {
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
