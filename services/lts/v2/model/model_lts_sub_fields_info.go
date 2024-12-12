package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LtsSubFieldsInfo json子字段信息
type LtsSubFieldsInfo struct {

	// 字段类型
	FieldType LtsSubFieldsInfoFieldType `json:"fieldType"`

	// 字段名称
	FieldName string `json:"fieldName"`

	// 是否大小写敏感
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	// 是否包含中文
	IncludeChinese *bool `json:"includeChinese,omitempty"`

	// 分词符
	Tokenizer *string `json:"tokenizer,omitempty"`

	// 是否快速分析
	QuickAnalysis *bool `json:"quickAnalysis,omitempty"`

	// 特殊分词符
	Ascii *[]string `json:"ascii,omitempty"`
}

func (o LtsSubFieldsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsSubFieldsInfo struct{}"
	}

	return strings.Join([]string{"LtsSubFieldsInfo", string(data)}, " ")
}

type LtsSubFieldsInfoFieldType struct {
	value string
}

type LtsSubFieldsInfoFieldTypeEnum struct {
	STRING LtsSubFieldsInfoFieldType
	LONG   LtsSubFieldsInfoFieldType
	FLOAT  LtsSubFieldsInfoFieldType
}

func GetLtsSubFieldsInfoFieldTypeEnum() LtsSubFieldsInfoFieldTypeEnum {
	return LtsSubFieldsInfoFieldTypeEnum{
		STRING: LtsSubFieldsInfoFieldType{
			value: "string",
		},
		LONG: LtsSubFieldsInfoFieldType{
			value: "long",
		},
		FLOAT: LtsSubFieldsInfoFieldType{
			value: "float",
		},
	}
}

func (c LtsSubFieldsInfoFieldType) Value() string {
	return c.value
}

func (c LtsSubFieldsInfoFieldType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LtsSubFieldsInfoFieldType) UnmarshalJSON(b []byte) error {
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
