package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateKeywordsAlarmRuleResponse struct {

	// 关键词告警id
	KeywordsAlarmRuleId *string `json:"keywords_alarm_rule_id,omitempty" xml:"keywords_alarm_rule_id"`

	// 关键词告警名称
	KeywordsAlarmRuleName *string `json:"keywords_alarm_rule_name,omitempty" xml:"keywords_alarm_rule_name"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription *string `json:"keywords_alarm_rule_description,omitempty" xml:"keywords_alarm_rule_description"`

	// 关键词详细信息
	KeywordsRequests *[]KeywordsRequest `json:"keywords_requests,omitempty" xml:"keywords_requests"`

	Frequency *Frequency `json:"frequency,omitempty" xml:"frequency"`

	// 告警级别
	KeywordsAlarmLevel *UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel `json:"keywords_alarm_level,omitempty" xml:"keywords_alarm_level"`

	// 是否发送
	KeywordsAlarmSend *bool `json:"keywords_alarm_send,omitempty" xml:"keywords_alarm_send"`

	// domainId
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 创建时间(毫秒时间戳)
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间(毫秒时间戳)
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`

	// 语言
	Language *string `json:"language,omitempty" xml:"language"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`

	// 主题信息
	Topics *[]Topics `json:"topics,omitempty" xml:"topics"`

	// 暂无
	ConditionExpression *string `json:"condition_expression,omitempty" xml:"condition_expression"`

	// 暂无
	IndexId        *string `json:"indexId,omitempty" xml:"indexId"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateKeywordsAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeywordsAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeywordsAlarmRuleResponse", string(data)}, " ")
}

type UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel struct {
	value string
}

type UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevelEnum struct {
	INFO     UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel
	MINOR    UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel
	MAJOR    UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel
	CRITICAL UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel
}

func GetUpdateKeywordsAlarmRuleResponseKeywordsAlarmLevelEnum() UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevelEnum {
	return UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevelEnum{
		INFO: UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel{
			value: "Info",
		},
		MINOR: UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel{
			value: "Minor",
		},
		MAJOR: UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel{
			value: "Major",
		},
		CRITICAL: UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel{
			value: "Critical",
		},
	}
}

func (c UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel) Value() string {
	return c.value
}

func (c UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel) UnmarshalJSON(b []byte) error {
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
