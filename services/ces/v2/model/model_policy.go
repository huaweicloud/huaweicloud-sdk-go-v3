package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type Policy struct {
	// 指标名称

	MetricName *string `json:"metric_name,omitempty"`
	// 周期

	Period *PolicyPeriod `json:"period,omitempty"`
	// 聚合方式

	Filter *string `json:"filter,omitempty"`
	// 阈值符号

	ComparisonOperator *string `json:"comparison_operator,omitempty"`
	// 阈值

	Value *float32 `json:"value,omitempty"`
	// 单位

	Unit *string `json:"unit,omitempty"`
	// 次数

	Count *int32 `json:"count,omitempty"`
	// 类型

	Type *string `json:"type,omitempty"`
	// 抑制方式

	SuppressDuration *PolicySuppressDuration `json:"suppress_duration,omitempty"`
	// 告警级别

	Level *int32 `json:"level,omitempty"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}

type PolicyPeriod struct {
	value int32
}

type PolicyPeriodEnum struct {
	E_0     PolicyPeriod
	E_1     PolicyPeriod
	E_300   PolicyPeriod
	E_1200  PolicyPeriod
	E_3600  PolicyPeriod
	E_14400 PolicyPeriod
	E_86400 PolicyPeriod
}

func GetPolicyPeriodEnum() PolicyPeriodEnum {
	return PolicyPeriodEnum{
		E_0: PolicyPeriod{
			value: 0,
		}, E_1: PolicyPeriod{
			value: 1,
		}, E_300: PolicyPeriod{
			value: 300,
		}, E_1200: PolicyPeriod{
			value: 1200,
		}, E_3600: PolicyPeriod{
			value: 3600,
		}, E_14400: PolicyPeriod{
			value: 14400,
		}, E_86400: PolicyPeriod{
			value: 86400,
		},
	}
}

func (c PolicyPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type PolicySuppressDuration struct {
	value int32
}

type PolicySuppressDurationEnum struct {
	E_0     PolicySuppressDuration
	E_300   PolicySuppressDuration
	E_600   PolicySuppressDuration
	E_900   PolicySuppressDuration
	E_1800  PolicySuppressDuration
	E_3600  PolicySuppressDuration
	E_10800 PolicySuppressDuration
	E_21600 PolicySuppressDuration
	E_43200 PolicySuppressDuration
	E_86400 PolicySuppressDuration
}

func GetPolicySuppressDurationEnum() PolicySuppressDurationEnum {
	return PolicySuppressDurationEnum{
		E_0: PolicySuppressDuration{
			value: 0,
		}, E_300: PolicySuppressDuration{
			value: 300,
		}, E_600: PolicySuppressDuration{
			value: 600,
		}, E_900: PolicySuppressDuration{
			value: 900,
		}, E_1800: PolicySuppressDuration{
			value: 1800,
		}, E_3600: PolicySuppressDuration{
			value: 3600,
		}, E_10800: PolicySuppressDuration{
			value: 10800,
		}, E_21600: PolicySuppressDuration{
			value: 21600,
		}, E_43200: PolicySuppressDuration{
			value: 43200,
		}, E_86400: PolicySuppressDuration{
			value: 86400,
		},
	}
}

func (c PolicySuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicySuppressDuration) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
