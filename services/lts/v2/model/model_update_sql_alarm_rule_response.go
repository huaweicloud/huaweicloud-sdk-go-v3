package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateSqlAlarmRuleResponse struct {
	// 测试

	Id *string `json:"id,omitempty"`
	// 测试

	IndexId *string `json:"indexId,omitempty"`
	// 测试

	Key *string `json:"key,omitempty"`
	// 测试

	Language *string `json:"language,omitempty"`
	// 测试

	ProjectId *string `json:"projectId,omitempty"`
	// SQL告警名称

	SqlAlarmRuleName *string `json:"sql_alarm_rule_name,omitempty"`
	// SQL告警规则id

	SqlAlarmRuleId *string `json:"sql_alarm_rule_id,omitempty"`
	// SQL告警信息描述

	SqlAlarmRuleDescription *string `json:"sql_alarm_rule_description,omitempty"`
	// SQL详细信息

	SqlRequests *[]SqlRequest `json:"sql_requests,omitempty"`
	// 告警统计周期

	Frequency *Frequency `json:"frequency,omitempty"`
	// 条件表达式

	ConditionExpression *string `json:"condition_expression,omitempty"`
	// 告警级别

	SqlAlarmLevel *UpdateSqlAlarmRuleResponseSqlAlarmLevel `json:"sql_alarm_level,omitempty"`
	// 是否发送

	SqlAlarmSend *bool `json:"sql_alarm_send,omitempty"`
	// domainId

	DomainId *string `json:"domain_id,omitempty"`
	// 创建时间(毫秒时间戳)

	CreateTime *interface{} `json:"create_time,omitempty"`
	// 更新时间(毫秒时间戳)

	UpdateTime *interface{} `json:"update_time,omitempty"`
	// 主题

	Topics         *[]Topics `json:"topics,omitempty"`
	HttpStatusCode int       `json:"-"`
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

func (c UpdateSqlAlarmRuleResponseSqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlAlarmRuleResponseSqlAlarmLevel) UnmarshalJSON(b []byte) error {
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
