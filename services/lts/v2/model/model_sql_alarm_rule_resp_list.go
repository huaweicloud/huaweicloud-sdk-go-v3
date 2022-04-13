package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

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

	TemplateName *string `json:"template_name,omitempty"`

	Status *SqlAlarmRuleRespListStatus `json:"status,omitempty"`
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

type SqlAlarmRuleRespListStatus struct {
	value string
}

type SqlAlarmRuleRespListStatusEnum struct {
	RUNNING  SqlAlarmRuleRespListStatus
	STOPPING SqlAlarmRuleRespListStatus
}

func GetSqlAlarmRuleRespListStatusEnum() SqlAlarmRuleRespListStatusEnum {
	return SqlAlarmRuleRespListStatusEnum{
		RUNNING: SqlAlarmRuleRespListStatus{
			value: "RUNNING",
		},
		STOPPING: SqlAlarmRuleRespListStatus{
			value: "STOPPING",
		},
	}
}

func (c SqlAlarmRuleRespListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlAlarmRuleRespListStatus) UnmarshalJSON(b []byte) error {
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
