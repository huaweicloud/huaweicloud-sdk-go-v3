package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KeyIndex struct {

	// 字段类型；text 全文索引字段、keyword 结构化字段、long Long、integer Integer、double Double、float Float、date 时间字段
	Type *KeyIndexType `json:"type,omitempty"`

	// 是否包含中文
	IsChineseExist *bool `json:"is_chinese_exist,omitempty"`

	// 嵌套结构
	Properties map[string]KeyIndex `json:"properties,omitempty"`
}

func (o KeyIndex) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyIndex struct{}"
	}

	return strings.Join([]string{"KeyIndex", string(data)}, " ")
}

type KeyIndexType struct {
	value string
}

type KeyIndexTypeEnum struct {
	TEXT    KeyIndexType
	KEYWORD KeyIndexType
	LONG    KeyIndexType
	INTEGER KeyIndexType
	DOUBLE  KeyIndexType
	FLOAT   KeyIndexType
	DATE    KeyIndexType
}

func GetKeyIndexTypeEnum() KeyIndexTypeEnum {
	return KeyIndexTypeEnum{
		TEXT: KeyIndexType{
			value: "text",
		},
		KEYWORD: KeyIndexType{
			value: "keyword",
		},
		LONG: KeyIndexType{
			value: "long",
		},
		INTEGER: KeyIndexType{
			value: "integer",
		},
		DOUBLE: KeyIndexType{
			value: "double",
		},
		FLOAT: KeyIndexType{
			value: "float",
		},
		DATE: KeyIndexType{
			value: "date",
		},
	}
}

func (c KeyIndexType) Value() string {
	return c.value
}

func (c KeyIndexType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeyIndexType) UnmarshalJSON(b []byte) error {
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
