package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRuleTrigger struct {

	// 累计次数
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// 表达式
	Expression string `json:"expression"`

	// **参数解释**: 指标模式 - COUNT 数量  **约束限制** 不涉及 **取值范围**: - COUNT  **默认值** 不涉及
	Mode *AlertRuleTriggerMode `json:"mode,omitempty"`

	// **参数解释**: 操作符类型 - GT 大于 - LT 小于 - EQ 等于 - NE 不等于  **约束限制** 不涉及 **取值范围**: - GT - LT - EQ - NE  **默认值** 不涉及
	Operator *AlertRuleTriggerOperator `json:"operator,omitempty"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity *AlertRuleTriggerSeverity `json:"severity,omitempty"`
}

func (o AlertRuleTrigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleTrigger struct{}"
	}

	return strings.Join([]string{"AlertRuleTrigger", string(data)}, " ")
}

type AlertRuleTriggerMode struct {
	value string
}

type AlertRuleTriggerModeEnum struct {
	COUNT AlertRuleTriggerMode
}

func GetAlertRuleTriggerModeEnum() AlertRuleTriggerModeEnum {
	return AlertRuleTriggerModeEnum{
		COUNT: AlertRuleTriggerMode{
			value: "COUNT",
		},
	}
}

func (c AlertRuleTriggerMode) Value() string {
	return c.value
}

func (c AlertRuleTriggerMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleTriggerMode) UnmarshalJSON(b []byte) error {
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

type AlertRuleTriggerOperator struct {
	value string
}

type AlertRuleTriggerOperatorEnum struct {
	EQ AlertRuleTriggerOperator
	NE AlertRuleTriggerOperator
	GT AlertRuleTriggerOperator
	LT AlertRuleTriggerOperator
}

func GetAlertRuleTriggerOperatorEnum() AlertRuleTriggerOperatorEnum {
	return AlertRuleTriggerOperatorEnum{
		EQ: AlertRuleTriggerOperator{
			value: "EQ",
		},
		NE: AlertRuleTriggerOperator{
			value: "NE",
		},
		GT: AlertRuleTriggerOperator{
			value: "GT",
		},
		LT: AlertRuleTriggerOperator{
			value: "LT",
		},
	}
}

func (c AlertRuleTriggerOperator) Value() string {
	return c.value
}

func (c AlertRuleTriggerOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleTriggerOperator) UnmarshalJSON(b []byte) error {
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

type AlertRuleTriggerSeverity struct {
	value string
}

type AlertRuleTriggerSeverityEnum struct {
	TIPS   AlertRuleTriggerSeverity
	LOW    AlertRuleTriggerSeverity
	MEDIUM AlertRuleTriggerSeverity
	HIGH   AlertRuleTriggerSeverity
	FATAL  AlertRuleTriggerSeverity
}

func GetAlertRuleTriggerSeverityEnum() AlertRuleTriggerSeverityEnum {
	return AlertRuleTriggerSeverityEnum{
		TIPS: AlertRuleTriggerSeverity{
			value: "TIPS",
		},
		LOW: AlertRuleTriggerSeverity{
			value: "LOW",
		},
		MEDIUM: AlertRuleTriggerSeverity{
			value: "MEDIUM",
		},
		HIGH: AlertRuleTriggerSeverity{
			value: "HIGH",
		},
		FATAL: AlertRuleTriggerSeverity{
			value: "FATAL",
		},
	}
}

func (c AlertRuleTriggerSeverity) Value() string {
	return c.value
}

func (c AlertRuleTriggerSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleTriggerSeverity) UnmarshalJSON(b []byte) error {
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
