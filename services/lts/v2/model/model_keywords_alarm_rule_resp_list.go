package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KeywordsAlarmRuleRespList struct {

	// 项目id
	ProjectId *string `json:"projectId,omitempty"`

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

	Frequency *Frequency `json:"frequency"`

	// 告警级别
	KeywordsAlarmLevel KeywordsAlarmRuleRespListKeywordsAlarmLevel `json:"keywords_alarm_level"`

	// 是否发送
	KeywordsAlarmSend bool `json:"keywords_alarm_send"`

	// domainId
	DomainId string `json:"domain_id"`

	// 创建时间（毫秒时间戳）
	CreateTime int64 `json:"create_time"`

	// 更新时间（毫秒时间戳）
	UpdateTime int64 `json:"update_time"`

	// 通知主题
	Topics []Topics `json:"topics"`

	// 消息模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 告警状态
	Status *KeywordsAlarmRuleRespListStatus `json:"status,omitempty"`

	// 触发条件：触发周期;默认为1
	TriggerConditionCount *int32 `json:"trigger_condition_count,omitempty"`

	// 触发条件：触发次数;默认为1
	TriggerConditionFrequency *int32 `json:"trigger_condition_frequency,omitempty"`

	// 是否打开恢复通知;默认false
	WhetherRecoveryPolicy *bool `json:"whether_recovery_policy,omitempty"`

	// 恢复策略周期;默认为3
	RecoveryPolicy *int32 `json:"recovery_policy,omitempty"`

	// 通知频率,单位(分钟)
	NotificationFrequency KeywordsAlarmRuleRespListNotificationFrequency `json:"notification_frequency"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`
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

type KeywordsAlarmRuleRespListStatus struct {
	value string
}

type KeywordsAlarmRuleRespListStatusEnum struct {
	RUNNING_  KeywordsAlarmRuleRespListStatus
	STOPPING_ KeywordsAlarmRuleRespListStatus
}

func GetKeywordsAlarmRuleRespListStatusEnum() KeywordsAlarmRuleRespListStatusEnum {
	return KeywordsAlarmRuleRespListStatusEnum{
		RUNNING_: KeywordsAlarmRuleRespListStatus{
			value: "RUNNING  启用",
		},
		STOPPING_: KeywordsAlarmRuleRespListStatus{
			value: "STOPPING  停止",
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

type KeywordsAlarmRuleRespListNotificationFrequency struct {
	value int32
}

type KeywordsAlarmRuleRespListNotificationFrequencyEnum struct {
	E_0   KeywordsAlarmRuleRespListNotificationFrequency
	E_5   KeywordsAlarmRuleRespListNotificationFrequency
	E_10  KeywordsAlarmRuleRespListNotificationFrequency
	E_15  KeywordsAlarmRuleRespListNotificationFrequency
	E_30  KeywordsAlarmRuleRespListNotificationFrequency
	E_60  KeywordsAlarmRuleRespListNotificationFrequency
	E_180 KeywordsAlarmRuleRespListNotificationFrequency
	E_360 KeywordsAlarmRuleRespListNotificationFrequency
}

func GetKeywordsAlarmRuleRespListNotificationFrequencyEnum() KeywordsAlarmRuleRespListNotificationFrequencyEnum {
	return KeywordsAlarmRuleRespListNotificationFrequencyEnum{
		E_0: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 0,
		}, E_5: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 5,
		}, E_10: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 10,
		}, E_15: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 15,
		}, E_30: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 30,
		}, E_60: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 60,
		}, E_180: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 180,
		}, E_360: KeywordsAlarmRuleRespListNotificationFrequency{
			value: 360,
		},
	}
}

func (c KeywordsAlarmRuleRespListNotificationFrequency) Value() int32 {
	return c.value
}

func (c KeywordsAlarmRuleRespListNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeywordsAlarmRuleRespListNotificationFrequency) UnmarshalJSON(b []byte) error {
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
