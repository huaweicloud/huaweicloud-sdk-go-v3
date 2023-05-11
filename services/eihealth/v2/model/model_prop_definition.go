package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 属性定义的相关结构
type PropDefinition struct {

	// 属性业务侧ID
	Id *string `json:"id,omitempty"`

	// 属性名称
	Name *string `json:"name,omitempty"`

	// 属性类型
	Type *PropDefinitionType `json:"type,omitempty"`

	// 属性具体描述信息
	Description *string `json:"description,omitempty"`

	ValueRange *ValueRange `json:"value_range,omitempty"`

	OptimalRange *ValueRange `json:"optimal_range,omitempty"`

	WarningRange *ValueRange `json:"warning_range,omitempty"`

	// 模型参数呈现类型
	Style *PropDefinitionStyle `json:"style,omitempty"`

	// 模型推理是否呈现置信区间
	ConfidentialInterval *bool `json:"confidential_interval,omitempty"`
}

func (o PropDefinition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropDefinition struct{}"
	}

	return strings.Join([]string{"PropDefinition", string(data)}, " ")
}

type PropDefinitionType struct {
	value string
}

type PropDefinitionTypeEnum struct {
	BINARY    PropDefinitionType
	NUMERICAL PropDefinitionType
}

func GetPropDefinitionTypeEnum() PropDefinitionTypeEnum {
	return PropDefinitionTypeEnum{
		BINARY: PropDefinitionType{
			value: "binary",
		},
		NUMERICAL: PropDefinitionType{
			value: "numerical",
		},
	}
}

func (c PropDefinitionType) Value() string {
	return c.value
}

func (c PropDefinitionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PropDefinitionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PropDefinitionStyle struct {
	value string
}

type PropDefinitionStyleEnum struct {
	NUMBER      PropDefinitionStyle
	PROBABILITY PropDefinitionStyle
}

func GetPropDefinitionStyleEnum() PropDefinitionStyleEnum {
	return PropDefinitionStyleEnum{
		NUMBER: PropDefinitionStyle{
			value: "number",
		},
		PROBABILITY: PropDefinitionStyle{
			value: "probability",
		},
	}
}

func (c PropDefinitionStyle) Value() string {
	return c.value
}

func (c PropDefinitionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PropDefinitionStyle) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
