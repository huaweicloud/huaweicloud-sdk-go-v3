package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TriggerPolicy struct {

	// 触发类型，accumulative: 累计触发，immediately: 立即触发。
	TriggerType TriggerPolicyTriggerType `json:"trigger_type"`

	// 触发周期，选择累计触发时需设置该参数，默认单位为s，支持5分钟、20分钟、1小时、4小时、24小时。
	Period *TriggerPolicyPeriod `json:"period,omitempty"`

	// 比较符，支持'>'和'>='。
	Operator *TriggerPolicyOperator `json:"operator,omitempty"`

	// 触发次数，选择累计触发时需设置该参数。
	Count *int32 `json:"count,omitempty"`
}

func (o TriggerPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerPolicy struct{}"
	}

	return strings.Join([]string{"TriggerPolicy", string(data)}, " ")
}

type TriggerPolicyTriggerType struct {
	value string
}

type TriggerPolicyTriggerTypeEnum struct {
	ACCUMULATIVE TriggerPolicyTriggerType
	IMMEDIATELY  TriggerPolicyTriggerType
}

func GetTriggerPolicyTriggerTypeEnum() TriggerPolicyTriggerTypeEnum {
	return TriggerPolicyTriggerTypeEnum{
		ACCUMULATIVE: TriggerPolicyTriggerType{
			value: "accumulative",
		},
		IMMEDIATELY: TriggerPolicyTriggerType{
			value: "immediately",
		},
	}
}

func (c TriggerPolicyTriggerType) Value() string {
	return c.value
}

func (c TriggerPolicyTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerPolicyTriggerType) UnmarshalJSON(b []byte) error {
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

type TriggerPolicyPeriod struct {
	value int32
}

type TriggerPolicyPeriodEnum struct {
	E_300   TriggerPolicyPeriod
	E_1200  TriggerPolicyPeriod
	E_3600  TriggerPolicyPeriod
	E_14400 TriggerPolicyPeriod
	E_86400 TriggerPolicyPeriod
}

func GetTriggerPolicyPeriodEnum() TriggerPolicyPeriodEnum {
	return TriggerPolicyPeriodEnum{
		E_300: TriggerPolicyPeriod{
			value: 300,
		}, E_1200: TriggerPolicyPeriod{
			value: 1200,
		}, E_3600: TriggerPolicyPeriod{
			value: 3600,
		}, E_14400: TriggerPolicyPeriod{
			value: 14400,
		}, E_86400: TriggerPolicyPeriod{
			value: 86400,
		},
	}
}

func (c TriggerPolicyPeriod) Value() int32 {
	return c.value
}

func (c TriggerPolicyPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerPolicyPeriod) UnmarshalJSON(b []byte) error {
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

type TriggerPolicyOperator struct {
	value string
}

type TriggerPolicyOperatorEnum struct {
	GREATER_THAN             TriggerPolicyOperator
	GREATER_THAN_OR_EQUAL_TO TriggerPolicyOperator
}

func GetTriggerPolicyOperatorEnum() TriggerPolicyOperatorEnum {
	return TriggerPolicyOperatorEnum{
		GREATER_THAN: TriggerPolicyOperator{
			value: ">",
		},
		GREATER_THAN_OR_EQUAL_TO: TriggerPolicyOperator{
			value: ">=",
		},
	}
}

func (c TriggerPolicyOperator) Value() string {
	return c.value
}

func (c TriggerPolicyOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerPolicyOperator) UnmarshalJSON(b []byte) error {
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
