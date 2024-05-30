package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ParaGroupParameterResult struct {

	// 特定参数名称。
	Name string `json:"name"`

	// 特定参数值。
	Value string `json:"value"`

	// 参数是否需要重启。 - 取值为\"true\"，需要重启。 - 取值为\"false\"，不需要重启。
	NeedRestart bool `json:"need_restart"`

	// 该参数是否只读(true：只读；false：可编辑)。
	Readonly bool `json:"readonly"`

	// 参数取值范围。
	ValueRange string `json:"value_range"`

	// 参数类型，取值为“string”、“integer”、“boolean”、“list”、\"all\"或“float”之一。
	DataType ParaGroupParameterResultDataType `json:"data_type"`

	// 参数描述。
	Description string `json:"description"`
}

func (o ParaGroupParameterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParaGroupParameterResult struct{}"
	}

	return strings.Join([]string{"ParaGroupParameterResult", string(data)}, " ")
}

type ParaGroupParameterResultDataType struct {
	value string
}

type ParaGroupParameterResultDataTypeEnum struct {
	STRING  ParaGroupParameterResultDataType
	INTEGER ParaGroupParameterResultDataType
	BOOLEAN ParaGroupParameterResultDataType
	LIST    ParaGroupParameterResultDataType
	FLOAT   ParaGroupParameterResultDataType
}

func GetParaGroupParameterResultDataTypeEnum() ParaGroupParameterResultDataTypeEnum {
	return ParaGroupParameterResultDataTypeEnum{
		STRING: ParaGroupParameterResultDataType{
			value: "string",
		},
		INTEGER: ParaGroupParameterResultDataType{
			value: "integer",
		},
		BOOLEAN: ParaGroupParameterResultDataType{
			value: "boolean",
		},
		LIST: ParaGroupParameterResultDataType{
			value: "list",
		},
		FLOAT: ParaGroupParameterResultDataType{
			value: "float",
		},
	}
}

func (c ParaGroupParameterResultDataType) Value() string {
	return c.value
}

func (c ParaGroupParameterResultDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParaGroupParameterResultDataType) UnmarshalJSON(b []byte) error {
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
