package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateKeywordsAlarmRuleRequestBody struct {

	// 关键词告警名称
	KeywordsAlarmRuleName string `json:"keywords_alarm_rule_name" xml:"keywords_alarm_rule_name"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription *string `json:"keywords_alarm_rule_description,omitempty" xml:"keywords_alarm_rule_description"`

	// 关键词详细信息
	KeywordsRequests []KeywordsRequest `json:"keywords_requests" xml:"keywords_requests"`

	Frequency *Frequency `json:"frequency" xml:"frequency"`

	// 告警级别
	KeywordsAlarmLevel CreateKeywordsAlarmRuleRequestBodyKeywordsAlarmLevel `json:"keywords_alarm_level" xml:"keywords_alarm_level"`

	// 是否发送
	KeywordsAlarmSend bool `json:"keywords_alarm_send" xml:"keywords_alarm_send"`

	// domainId
	DomainId string `json:"domain_id" xml:"domain_id"`

	// 通知主题
	NotificationSaveRule *NotificationSaveRule `json:"notification_save_rule,omitempty" xml:"notification_save_rule"`
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
			value: "CRITICAL",
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
