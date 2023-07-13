package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateKeywordsAlarmRuleResponse Response Object
type UpdateKeywordsAlarmRuleResponse struct {

	// 关键词告警id
	KeywordsAlarmRuleId *string `json:"keywords_alarm_rule_id,omitempty"`

	// 关键词告警名称
	KeywordsAlarmRuleName *string `json:"keywords_alarm_rule_name,omitempty"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription *string `json:"keywords_alarm_rule_description,omitempty"`

	// 关键词详细信息
	KeywordsRequests *[]KeywordsRequest `json:"keywords_requests,omitempty"`

	Frequency *Frequency `json:"frequency,omitempty"`

	// 告警级别
	KeywordsAlarmLevel *UpdateKeywordsAlarmRuleResponseKeywordsAlarmLevel `json:"keywords_alarm_level,omitempty"`

	// 是否发送
	KeywordsAlarmSend *bool `json:"keywords_alarm_send,omitempty"`

	// domainId
	DomainId *string `json:"domain_id,omitempty"`

	// 创建时间(毫秒时间戳)
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间(毫秒时间戳)
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty"`

	// 主题信息
	Topics *[]Topics `json:"topics,omitempty"`

	// 暂无
	ConditionExpression *string `json:"condition_expression,omitempty"`

	// 暂无
	IndexId        *string `json:"indexId,omitempty"`
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
