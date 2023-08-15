package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

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
