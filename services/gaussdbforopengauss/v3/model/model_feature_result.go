package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FeatureResult struct {

	// 特性名称。
	Name *string `json:"name,omitempty"`

	// 特性是否开启。
	Status *string `json:"status,omitempty"`

	// 特性中文描述。
	Description *string `json:"description,omitempty"`

	// 特性值类型。
	Type *FeatureResultType `json:"type,omitempty"`

	// 特性值。
	Value *string `json:"value,omitempty"`

	// 特性值范围。
	Range *string `json:"range,omitempty"`

	// 特性范围描述。
	RangeDescription *string `json:"range_description,omitempty"`
}

func (o FeatureResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureResult struct{}"
	}

	return strings.Join([]string{"FeatureResult", string(data)}, " ")
}

type FeatureResultType struct {
	value string
}

type FeatureResultTypeEnum struct {
	INTEGER FeatureResultType
	STRING  FeatureResultType
	BOOLEAN FeatureResultType
}

func GetFeatureResultTypeEnum() FeatureResultTypeEnum {
	return FeatureResultTypeEnum{
		INTEGER: FeatureResultType{
			value: "integer",
		},
		STRING: FeatureResultType{
			value: "string",
		},
		BOOLEAN: FeatureResultType{
			value: "boolean",
		},
	}
}

func (c FeatureResultType) Value() string {
	return c.value
}

func (c FeatureResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FeatureResultType) UnmarshalJSON(b []byte) error {
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
