package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateSqlAlarmRuleResponse struct {

	// 测试
	Id *string `json:"id,omitempty" xml:"id"`

	// 测试
	IndexId *string `json:"indexId,omitempty" xml:"indexId"`

	// 测试
	Language *string `json:"language,omitempty" xml:"language"`

	// 测试
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`

	// SQL告警名称
	SqlAlarmRuleName *string `json:"sql_alarm_rule_name,omitempty" xml:"sql_alarm_rule_name"`

	// SQL告警规则id
	SqlAlarmRuleId *string `json:"sql_alarm_rule_id,omitempty" xml:"sql_alarm_rule_id"`

	// SQL告警信息描述
	SqlAlarmRuleDescription *string `json:"sql_alarm_rule_description,omitempty" xml:"sql_alarm_rule_description"`

	// SQL详细信息
	SqlRequests *[]SqlRequest `json:"sql_requests,omitempty" xml:"sql_requests"`

	// 告警统计周期
	Frequency *Frequency `json:"frequency,omitempty" xml:"frequency"`

	// 条件表达式
	ConditionExpression *string `json:"condition_expression,omitempty" xml:"condition_expression"`

	// 告警级别
	SqlAlarmLevel *UpdateSqlAlarmRuleResponseSqlAlarmLevel `json:"sql_alarm_level,omitempty" xml:"sql_alarm_level"`

	// 是否发送
	SqlAlarmSend *bool `json:"sql_alarm_send,omitempty" xml:"sql_alarm_send"`

	// domainId
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 创建时间(毫秒时间戳)
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间(毫秒时间戳)
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`

	// 主题
	Topics         *[]Topics `json:"topics,omitempty" xml:"topics"`
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

func (c UpdateSqlAlarmRuleResponseSqlAlarmLevel) Value() string {
	return c.value
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
