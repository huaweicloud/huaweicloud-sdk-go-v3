package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSlaCustomizedTemplateResponse Response Object
type ShowSlaCustomizedTemplateResponse struct {

	// Id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 触发类型(EVENT_TICKET,ALARM_TICKET,CHANGE_NOTE,TO_DO_TASKS,ISSUE_TICKET)
	TriggerType *ShowSlaCustomizedTemplateResponseTriggerType `json:"trigger_type,omitempty"`

	// 触发等级
	TriggerLevel *string `json:"trigger_level,omitempty"`

	// 应用
	Application *string `json:"application,omitempty"`

	// 默认模板
	DefaultTemplate *bool `json:"default_template,omitempty"`

	// 启用状态
	Enabled *bool `json:"enabled,omitempty"`

	// 规则信息
	SlaRules *[]SlaRuleInfo `json:"sla_rules,omitempty"`

	// 生效类型（MON_SUN_24_HOURS一直生效、SPECIFIC_TIME阶段性生效）
	EffectiveType *ShowSlaCustomizedTemplateResponseEffectiveType `json:"effective_type,omitempty"`

	// 生效周期（每天、周一、周二、周三、周四、周五、周六、周日）
	EffectivePeriod *string `json:"effective_period,omitempty"`

	// 生效开始时间
	EffectiveStartTime *string `json:"effective_start_time,omitempty"`

	// 生效结束时间
	EffectiveEndTime *string `json:"effective_end_time,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowSlaCustomizedTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlaCustomizedTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowSlaCustomizedTemplateResponse", string(data)}, " ")
}

type ShowSlaCustomizedTemplateResponseTriggerType struct {
	value string
}

type ShowSlaCustomizedTemplateResponseTriggerTypeEnum struct {
	EVENT_TICKET ShowSlaCustomizedTemplateResponseTriggerType
	ALARM_TICKET ShowSlaCustomizedTemplateResponseTriggerType
	CHANGE_NOTE  ShowSlaCustomizedTemplateResponseTriggerType
	TO_DO_TASKS  ShowSlaCustomizedTemplateResponseTriggerType
	ISSUE_TICKET ShowSlaCustomizedTemplateResponseTriggerType
}

func GetShowSlaCustomizedTemplateResponseTriggerTypeEnum() ShowSlaCustomizedTemplateResponseTriggerTypeEnum {
	return ShowSlaCustomizedTemplateResponseTriggerTypeEnum{
		EVENT_TICKET: ShowSlaCustomizedTemplateResponseTriggerType{
			value: "EVENT_TICKET",
		},
		ALARM_TICKET: ShowSlaCustomizedTemplateResponseTriggerType{
			value: "ALARM_TICKET",
		},
		CHANGE_NOTE: ShowSlaCustomizedTemplateResponseTriggerType{
			value: "CHANGE_NOTE",
		},
		TO_DO_TASKS: ShowSlaCustomizedTemplateResponseTriggerType{
			value: "TO_DO_TASKS",
		},
		ISSUE_TICKET: ShowSlaCustomizedTemplateResponseTriggerType{
			value: "ISSUE_TICKET",
		},
	}
}

func (c ShowSlaCustomizedTemplateResponseTriggerType) Value() string {
	return c.value
}

func (c ShowSlaCustomizedTemplateResponseTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSlaCustomizedTemplateResponseTriggerType) UnmarshalJSON(b []byte) error {
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

type ShowSlaCustomizedTemplateResponseEffectiveType struct {
	value string
}

type ShowSlaCustomizedTemplateResponseEffectiveTypeEnum struct {
	MON_SUN_24_HOURS ShowSlaCustomizedTemplateResponseEffectiveType
	SPECIFIC_TIME    ShowSlaCustomizedTemplateResponseEffectiveType
}

func GetShowSlaCustomizedTemplateResponseEffectiveTypeEnum() ShowSlaCustomizedTemplateResponseEffectiveTypeEnum {
	return ShowSlaCustomizedTemplateResponseEffectiveTypeEnum{
		MON_SUN_24_HOURS: ShowSlaCustomizedTemplateResponseEffectiveType{
			value: "MON_SUN_24_HOURS",
		},
		SPECIFIC_TIME: ShowSlaCustomizedTemplateResponseEffectiveType{
			value: "SPECIFIC_TIME",
		},
	}
}

func (c ShowSlaCustomizedTemplateResponseEffectiveType) Value() string {
	return c.value
}

func (c ShowSlaCustomizedTemplateResponseEffectiveType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSlaCustomizedTemplateResponseEffectiveType) UnmarshalJSON(b []byte) error {
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
