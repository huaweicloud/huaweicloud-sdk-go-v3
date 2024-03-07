package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateParameterDefinition 预定义合规包模板参数详情。
type TemplateParameterDefinition struct {

	// 预定义合规包模板参数名字。
	Name *string `json:"name,omitempty"`

	// 预定义合规包模板参数描述。
	Description *string `json:"description,omitempty"`

	// 预定义合规包模板参数默认值。
	DefaultValue *interface{} `json:"default_value,omitempty"`

	// 预定义合规包模板参数允许值列表。
	AllowedValues *[]interface{} `json:"allowed_values,omitempty"`

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

	// 预定义合规包模板参数类型。
	Type *TemplateParameterDefinitionType `json:"type,omitempty"`
}

func (o TemplateParameterDefinition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateParameterDefinition struct{}"
	}

	return strings.Join([]string{"TemplateParameterDefinition", string(data)}, " ")
}

type TemplateParameterDefinitionType struct {
	value string
}

type TemplateParameterDefinitionTypeEnum struct {
	ARRAY   TemplateParameterDefinitionType
	BOOLEAN TemplateParameterDefinitionType
	INTEGER TemplateParameterDefinitionType
	FLOAT   TemplateParameterDefinitionType
	STRING  TemplateParameterDefinitionType
	OBJECT  TemplateParameterDefinitionType
	DATE    TemplateParameterDefinitionType
}

func GetTemplateParameterDefinitionTypeEnum() TemplateParameterDefinitionTypeEnum {
	return TemplateParameterDefinitionTypeEnum{
		ARRAY: TemplateParameterDefinitionType{
			value: "Array",
		},
		BOOLEAN: TemplateParameterDefinitionType{
			value: "Boolean",
		},
		INTEGER: TemplateParameterDefinitionType{
			value: "Integer",
		},
		FLOAT: TemplateParameterDefinitionType{
			value: "Float",
		},
		STRING: TemplateParameterDefinitionType{
			value: "String",
		},
		OBJECT: TemplateParameterDefinitionType{
			value: "Object",
		},
		DATE: TemplateParameterDefinitionType{
			value: "Date",
		},
	}
}

func (c TemplateParameterDefinitionType) Value() string {
	return c.value
}

func (c TemplateParameterDefinitionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateParameterDefinitionType) UnmarshalJSON(b []byte) error {
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
