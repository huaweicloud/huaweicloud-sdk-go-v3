package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyParameterDefinition 策略参数定义
type PolicyParameterDefinition struct {

	// 策略参数名字
	Name *string `json:"name,omitempty"`

	// 策略参数描述
	Description *string `json:"description,omitempty"`

	// 策略参数允许值列表
	AllowedValues *[]interface{} `json:"allowed_values,omitempty"`

	// 策略参数默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 策略参数的最小值，当参数类型为Integer或Float时生效。
	Minimum *float32 `json:"minimum,omitempty"`

	// 策略参数的最大值，当参数类型为Integer或Float时生效。
	Maximum *float32 `json:"maximum,omitempty"`

	// 策略参数的最小项数，当参数类型为Array时生效。
	MinItems *int32 `json:"min_items,omitempty"`

	// 策略参数的最大项数，当参数类型为Array时生效。
	MaxItems *int32 `json:"max_items,omitempty"`

	// 策略参数的最小字符串长度或每项的最小字符串长度，当参数类型为String或Array时生效。
	MinLength *int32 `json:"min_length,omitempty"`

	// 策略参数的最大字符串长度或每项的最大字符串长度，当参数类型为String或Array时生效。
	MaxLength *int32 `json:"max_length,omitempty"`

	// 策略参数的字符串正则要求或每项的字符串正则要求，当参数类型为String或Array时生效。
	Pattern *string `json:"pattern,omitempty"`

	// 策略参数类型
	Type *PolicyParameterDefinitionType `json:"type,omitempty"`
}

func (o PolicyParameterDefinition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyParameterDefinition struct{}"
	}

	return strings.Join([]string{"PolicyParameterDefinition", string(data)}, " ")
}

type PolicyParameterDefinitionType struct {
	value string
}

type PolicyParameterDefinitionTypeEnum struct {
	ARRAY     PolicyParameterDefinitionType
	BOOLEAN   PolicyParameterDefinitionType
	INTEGER   PolicyParameterDefinitionType
	FLOAT     PolicyParameterDefinitionType
	STRING    PolicyParameterDefinitionType
	DATE_TIME PolicyParameterDefinitionType
}

func GetPolicyParameterDefinitionTypeEnum() PolicyParameterDefinitionTypeEnum {
	return PolicyParameterDefinitionTypeEnum{
		ARRAY: PolicyParameterDefinitionType{
			value: "Array",
		},
		BOOLEAN: PolicyParameterDefinitionType{
			value: "Boolean",
		},
		INTEGER: PolicyParameterDefinitionType{
			value: "Integer",
		},
		FLOAT: PolicyParameterDefinitionType{
			value: "Float",
		},
		STRING: PolicyParameterDefinitionType{
			value: "String",
		},
		DATE_TIME: PolicyParameterDefinitionType{
			value: "DateTime",
		},
	}
}

func (c PolicyParameterDefinitionType) Value() string {
	return c.value
}

func (c PolicyParameterDefinitionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyParameterDefinitionType) UnmarshalJSON(b []byte) error {
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
