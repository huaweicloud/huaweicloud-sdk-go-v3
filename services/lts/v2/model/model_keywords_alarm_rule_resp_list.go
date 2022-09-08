package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KeywordsAlarmRuleRespList struct {

	// 项目id
	ProjectId string `json:"projectId"`

	// 关键词告警id
	KeywordsAlarmRuleId string `json:"keywords_alarm_rule_id"`

	// 关键词告警名称
	KeywordsAlarmRuleName string `json:"keywords_alarm_rule_name"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription string `json:"keywords_alarm_rule_description"`

	// 条件
	ConditionExpression string `json:"condition_expression"`

	// 关键词详细信息
	KeywordsRequests []KeywordsRequest `json:"keywords_requests"`

	// 告警统计周期
	Frequency *Frequency `json:"frequency"`

	// 告警级别
	KeywordsAlarmLevel KeywordsAlarmRuleRespListKeywordsAlarmLevel `json:"keywords_alarm_level"`

	// 是否发送
	KeywordsAlarmSend bool `json:"keywords_alarm_send"`

	// domainId
	DomainId string `json:"domain_id"`

	// 创建时间(毫秒时间戳)
	CreateTime int64 `json:"create_time"`

	// 更新时间(毫秒时间戳)
	UpdateTime int64 `json:"update_time"`

	// 主题
	Topics []Topics `json:"topics"`

	TemplateName *string `json:"template_name,omitempty"`

	Status *KeywordsAlarmRuleRespListStatus `json:"status,omitempty"`
}

func (o KeywordsAlarmRuleRespList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeywordsAlarmRuleRespList struct{}"
	}

	return strings.Join([]string{"KeywordsAlarmRuleRespList", string(data)}, " ")
}

type KeywordsAlarmRuleRespListKeywordsAlarmLevel struct {
	value string
}

type KeywordsAlarmRuleRespListKeywordsAlarmLevelEnum struct {
	INFO     KeywordsAlarmRuleRespListKeywordsAlarmLevel
	MINOR    KeywordsAlarmRuleRespListKeywordsAlarmLevel
	MAJOR    KeywordsAlarmRuleRespListKeywordsAlarmLevel
	CRITICAL KeywordsAlarmRuleRespListKeywordsAlarmLevel
}

func GetKeywordsAlarmRuleRespListKeywordsAlarmLevelEnum() KeywordsAlarmRuleRespListKeywordsAlarmLevelEnum {
	return KeywordsAlarmRuleRespListKeywordsAlarmLevelEnum{
		INFO: KeywordsAlarmRuleRespListKeywordsAlarmLevel{
			value: "Info",
		},
		MINOR: KeywordsAlarmRuleRespListKeywordsAlarmLevel{
			value: "Minor",
		},
		MAJOR: KeywordsAlarmRuleRespListKeywordsAlarmLevel{
			value: "Major",
		},
		CRITICAL: KeywordsAlarmRuleRespListKeywordsAlarmLevel{
			value: "Critical",
		},
	}
}

func (c KeywordsAlarmRuleRespListKeywordsAlarmLevel) Value() string {
	return c.value
}

func (c KeywordsAlarmRuleRespListKeywordsAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeywordsAlarmRuleRespListKeywordsAlarmLevel) UnmarshalJSON(b []byte) error {
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

type KeywordsAlarmRuleRespListStatus struct {
	value string
}

type KeywordsAlarmRuleRespListStatusEnum struct {
	RUNNING  KeywordsAlarmRuleRespListStatus
	STOPPING KeywordsAlarmRuleRespListStatus
}

func GetKeywordsAlarmRuleRespListStatusEnum() KeywordsAlarmRuleRespListStatusEnum {
	return KeywordsAlarmRuleRespListStatusEnum{
		RUNNING: KeywordsAlarmRuleRespListStatus{
			value: "RUNNING",
		},
		STOPPING: KeywordsAlarmRuleRespListStatus{
			value: "STOPPING",
		},
	}
}

func (c KeywordsAlarmRuleRespListStatus) Value() string {
	return c.value
}

func (c KeywordsAlarmRuleRespListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeywordsAlarmRuleRespListStatus) UnmarshalJSON(b []byte) error {
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
