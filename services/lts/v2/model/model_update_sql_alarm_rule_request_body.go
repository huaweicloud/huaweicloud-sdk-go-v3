package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateSqlAlarmRuleRequestBody struct {
	// SQL告警id

	SqlAlarmRuleId string `json:"sql_alarm_rule_id"`
	// SQL告警名称

	SqlAlarmRuleName string `json:"sql_alarm_rule_name"`
	// SQL告警信息描述

	SqlAlarmRuleDescription *string `json:"sql_alarm_rule_description,omitempty"`
	// SQL详细信息

	SqlRequests []SqlRequest `json:"sql_requests"`
	// 告警统计周期

	Frequency *Frequency `json:"frequency"`
	// 条件表达式

	ConditionExpression string `json:"condition_expression"`
	// 告警级别

	SqlAlarmLevel UpdateSqlAlarmRuleRequestBodySqlAlarmLevel `json:"sql_alarm_level"`
	// 是否发送

	SqlAlarmSend bool `json:"sql_alarm_send"`
	// 发送主题 0:不变 1:新增 2:修改 3:删除

	SqlAlarmSendCode UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode `json:"sql_alarm_send_code"`
	// domainId

	DomainId string `json:"domain_id"`
	// 通知主题

	NotificationSaveRule *NotificationSaveRule `json:"notification_save_rule,omitempty"`
	// 邮件附加信息是否英文

	WhetherEnglish bool `json:"whether_english"`
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

func (c UpdateSqlAlarmRuleRequestBodySqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlAlarmRuleRequestBodySqlAlarmLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode struct {
	value int32
}

type UpdateSqlAlarmRuleRequestBodySqlAlarmSendCodeEnum struct {
	E_0 UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode
	E_1 UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode
	E_2 UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode
	E_3 UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode
}

func GetUpdateSqlAlarmRuleRequestBodySqlAlarmSendCodeEnum() UpdateSqlAlarmRuleRequestBodySqlAlarmSendCodeEnum {
	return UpdateSqlAlarmRuleRequestBodySqlAlarmSendCodeEnum{
		E_0: UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode{
			value: 0,
		}, E_1: UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode{
			value: 1,
		}, E_2: UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode{
			value: 2,
		}, E_3: UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode{
			value: 3,
		},
	}
}

func (c UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlAlarmRuleRequestBodySqlAlarmSendCode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
