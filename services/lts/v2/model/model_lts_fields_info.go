package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LtsFieldsInfo struct {

	// 字段类型
	FieldType LtsFieldsInfoFieldType `json:"fieldType"`

	// 字段名称
	FieldName string `json:"fieldName"`

	// 是否大小写敏感
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	// 是否包含中文
	IncludeChinese *bool `json:"includeChinese,omitempty"`

	// 分词符
	Tokenizer string `json:"tokenizer"`

	// 是否快速分析
	QuickAnalysis *bool `json:"quickAnalysis,omitempty"`

	// 特殊分词符
	Ascii *[]string `json:"ascii,omitempty"`

	// json字段信息
	LtsSubFieldsInfoList *[]LtsSubFieldsInfo `json:"ltsSubFieldsInfoList,omitempty"`
}

func (o LtsFieldsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsFieldsInfo struct{}"
	}

	return strings.Join([]string{"LtsFieldsInfo", string(data)}, " ")
}

type LtsFieldsInfoFieldType struct {
	value string
}

type LtsFieldsInfoFieldTypeEnum struct {
	STRING LtsFieldsInfoFieldType
	LONG   LtsFieldsInfoFieldType
	FLOAT  LtsFieldsInfoFieldType
	JSON   LtsFieldsInfoFieldType
}

func GetLtsFieldsInfoFieldTypeEnum() LtsFieldsInfoFieldTypeEnum {
	return LtsFieldsInfoFieldTypeEnum{
		STRING: LtsFieldsInfoFieldType{
			value: "string",
		},
		LONG: LtsFieldsInfoFieldType{
			value: "long",
		},
		FLOAT: LtsFieldsInfoFieldType{
			value: "float",
		},
		JSON: LtsFieldsInfoFieldType{
			value: "json",
		},
	}
}

func (c LtsFieldsInfoFieldType) Value() string {
	return c.value
}

func (c LtsFieldsInfoFieldType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LtsFieldsInfoFieldType) UnmarshalJSON(b []byte) error {
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
