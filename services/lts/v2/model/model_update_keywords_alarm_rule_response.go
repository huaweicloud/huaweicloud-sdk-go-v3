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

	// 规则名称
	AlarmRuleAlias *string `json:"alarm_rule_alias,omitempty"`

	// 关键词告警信息描述
	KeywordsAlarmRuleDescription *string `json:"keywords_alarm_rule_description,omitempty"`

	// 关键词详细信息
	KeywordsRequests *[]KeywordsResBody `json:"keywords_requests,omitempty"`

	Frequency *FrequencyRespBody `json:"frequency,omitempty"`

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

	// 邮件附加信息语言
	Language *UpdateKeywordsAlarmRuleResponseLanguage `json:"language,omitempty"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty"`

	// 通知主题
	Topics *[]Topics `json:"topics,omitempty"`

	// 情况表述
	ConditionExpression *string `json:"condition_expression,omitempty"`

	// 索引id
	IndexId *string `json:"indexId,omitempty"`

	// 通知频率,单位(分钟)
	NotificationFrequency *UpdateKeywordsAlarmRuleResponseNotificationFrequency `json:"notification_frequency,omitempty"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`
	HttpStatusCode      int     `json:"-"`
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

type UpdateKeywordsAlarmRuleResponseLanguage struct {
	value string
}

type UpdateKeywordsAlarmRuleResponseLanguageEnum struct {
	ZH_CN UpdateKeywordsAlarmRuleResponseLanguage
	EN_US UpdateKeywordsAlarmRuleResponseLanguage
}

func GetUpdateKeywordsAlarmRuleResponseLanguageEnum() UpdateKeywordsAlarmRuleResponseLanguageEnum {
	return UpdateKeywordsAlarmRuleResponseLanguageEnum{
		ZH_CN: UpdateKeywordsAlarmRuleResponseLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateKeywordsAlarmRuleResponseLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateKeywordsAlarmRuleResponseLanguage) Value() string {
	return c.value
}

func (c UpdateKeywordsAlarmRuleResponseLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKeywordsAlarmRuleResponseLanguage) UnmarshalJSON(b []byte) error {
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

type UpdateKeywordsAlarmRuleResponseNotificationFrequency struct {
	value int32
}

type UpdateKeywordsAlarmRuleResponseNotificationFrequencyEnum struct {
	E_0   UpdateKeywordsAlarmRuleResponseNotificationFrequency
	E_5   UpdateKeywordsAlarmRuleResponseNotificationFrequency
	E_10  UpdateKeywordsAlarmRuleResponseNotificationFrequency
	E_15  UpdateKeywordsAlarmRuleResponseNotificationFrequency
	E_30  UpdateKeywordsAlarmRuleResponseNotificationFrequency
	E_60  UpdateKeywordsAlarmRuleResponseNotificationFrequency
	E_180 UpdateKeywordsAlarmRuleResponseNotificationFrequency
	E_360 UpdateKeywordsAlarmRuleResponseNotificationFrequency
}

func GetUpdateKeywordsAlarmRuleResponseNotificationFrequencyEnum() UpdateKeywordsAlarmRuleResponseNotificationFrequencyEnum {
	return UpdateKeywordsAlarmRuleResponseNotificationFrequencyEnum{
		E_0: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 0,
		}, E_5: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 5,
		}, E_10: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 10,
		}, E_15: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 15,
		}, E_30: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 30,
		}, E_60: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 60,
		}, E_180: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 180,
		}, E_360: UpdateKeywordsAlarmRuleResponseNotificationFrequency{
			value: 360,
		},
	}
}

func (c UpdateKeywordsAlarmRuleResponseNotificationFrequency) Value() int32 {
	return c.value
}

func (c UpdateKeywordsAlarmRuleResponseNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKeywordsAlarmRuleResponseNotificationFrequency) UnmarshalJSON(b []byte) error {
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
