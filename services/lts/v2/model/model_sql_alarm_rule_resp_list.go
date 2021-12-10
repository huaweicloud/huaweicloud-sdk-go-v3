package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SqlAlarmRuleRespList struct {
	// SQL告警名称

	SqlAlarmRuleName string `json:"sql_alarm_rule_name"`
	// SQL告警规则id

	SqlAlarmRuleId string `json:"sql_alarm_rule_id"`
	// SQL告警信息描述

	SqlAlarmRuleDescription string `json:"sql_alarm_rule_description"`
	// SQL详细信息

	SqlRequests []SqlRequest `json:"sql_requests"`
	// 告警统计周期

	Frequency *Frequency `json:"frequency"`
	// 条件表达式

	ConditionExpression string `json:"condition_expression"`
	// 主题信息

	Topics []Topics `json:"topics"`
	// 告警级别

	SqlAlarmLevel SqlAlarmRuleRespListSqlAlarmLevel `json:"sql_alarm_level"`
	// 是否发送

	SqlAlarmSend bool `json:"sql_alarm_send"`
	// domainId

	DomainId string `json:"domain_id"`
	// 创建时间(毫秒时间戳)

	CreateTime int64 `json:"create_time"`
	// 更新时间(毫秒时间戳)

	UpdateTime int64 `json:"update_time"`
	// 邮件附加信息是否英文

	Language SqlAlarmRuleRespListLanguage `json:"language"`
}

func (o SqlAlarmRuleRespList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlAlarmRuleRespList struct{}"
	}

	return strings.Join([]string{"SqlAlarmRuleRespList", string(data)}, " ")
}

type SqlAlarmRuleRespListSqlAlarmLevel struct {
	value string
}

type SqlAlarmRuleRespListSqlAlarmLevelEnum struct {
	INFO     SqlAlarmRuleRespListSqlAlarmLevel
	MINOR    SqlAlarmRuleRespListSqlAlarmLevel
	MAJOR    SqlAlarmRuleRespListSqlAlarmLevel
	CRITICAL SqlAlarmRuleRespListSqlAlarmLevel
}

func GetSqlAlarmRuleRespListSqlAlarmLevelEnum() SqlAlarmRuleRespListSqlAlarmLevelEnum {
	return SqlAlarmRuleRespListSqlAlarmLevelEnum{
		INFO: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Info",
		},
		MINOR: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Minor",
		},
		MAJOR: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Major",
		},
		CRITICAL: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Critical",
		},
	}
}

func (c SqlAlarmRuleRespListSqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlAlarmRuleRespListSqlAlarmLevel) UnmarshalJSON(b []byte) error {
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

type SqlAlarmRuleRespListLanguage struct {
	value string
}

type SqlAlarmRuleRespListLanguageEnum struct {
	ZH_CN SqlAlarmRuleRespListLanguage
	EN_US SqlAlarmRuleRespListLanguage
}

func GetSqlAlarmRuleRespListLanguageEnum() SqlAlarmRuleRespListLanguageEnum {
	return SqlAlarmRuleRespListLanguageEnum{
		ZH_CN: SqlAlarmRuleRespListLanguage{
			value: "zh-cn",
		},
		EN_US: SqlAlarmRuleRespListLanguage{
			value: "en-us",
		},
	}
}

func (c SqlAlarmRuleRespListLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlAlarmRuleRespListLanguage) UnmarshalJSON(b []byte) error {
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
