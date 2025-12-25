package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AnalysisField struct {

	// 字段别名
	Alias *string `json:"alias,omitempty"`

	// 字段名称
	Name string `json:"name"`

	// 字段类型；可选值：boolean、byte、short、integer、long、float、half_float、scaled_float、double、keyword、text、date、ip、binary、object、nested
	Type AnalysisFieldType `json:"type"`
}

func (o AnalysisField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisField struct{}"
	}

	return strings.Join([]string{"AnalysisField", string(data)}, " ")
}

type AnalysisFieldType struct {
	value string
}

type AnalysisFieldTypeEnum struct {
	BOOLEAN      AnalysisFieldType
	BYTE         AnalysisFieldType
	SHORT        AnalysisFieldType
	INTEGER      AnalysisFieldType
	LONG         AnalysisFieldType
	FLOAT        AnalysisFieldType
	HALF_FLOAT   AnalysisFieldType
	SCALED_FLOAT AnalysisFieldType
	DOUBLE       AnalysisFieldType
	KEYWORD      AnalysisFieldType
	TEXT         AnalysisFieldType
	DATE         AnalysisFieldType
	IP           AnalysisFieldType
	BINARY       AnalysisFieldType
	OBJECT       AnalysisFieldType
	NESTED       AnalysisFieldType
}

func GetAnalysisFieldTypeEnum() AnalysisFieldTypeEnum {
	return AnalysisFieldTypeEnum{
		BOOLEAN: AnalysisFieldType{
			value: "boolean",
		},
		BYTE: AnalysisFieldType{
			value: "byte",
		},
		SHORT: AnalysisFieldType{
			value: "short",
		},
		INTEGER: AnalysisFieldType{
			value: "integer",
		},
		LONG: AnalysisFieldType{
			value: "long",
		},
		FLOAT: AnalysisFieldType{
			value: "float",
		},
		HALF_FLOAT: AnalysisFieldType{
			value: "half_float",
		},
		SCALED_FLOAT: AnalysisFieldType{
			value: "scaled_float",
		},
		DOUBLE: AnalysisFieldType{
			value: "double",
		},
		KEYWORD: AnalysisFieldType{
			value: "keyword",
		},
		TEXT: AnalysisFieldType{
			value: "text",
		},
		DATE: AnalysisFieldType{
			value: "date",
		},
		IP: AnalysisFieldType{
			value: "ip",
		},
		BINARY: AnalysisFieldType{
			value: "binary",
		},
		OBJECT: AnalysisFieldType{
			value: "object",
		},
		NESTED: AnalysisFieldType{
			value: "nested",
		},
	}
}

func (c AnalysisFieldType) Value() string {
	return c.value
}

func (c AnalysisFieldType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnalysisFieldType) UnmarshalJSON(b []byte) error {
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
