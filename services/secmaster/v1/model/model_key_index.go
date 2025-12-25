package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KeyIndex struct {

	// 是否包含中文
	IsChineseExist *bool `json:"is_chinese_exist,omitempty"`

	// 嵌套结构
	Properties map[string]KeyIndex `json:"properties,omitempty"`

	// **参数解释**: 字段类型 - text 全文索引字段 - keyword 关键字类型，适用于精确匹配 - long 长整型 - integer 整型 - double 双精度浮点数 - float 单精度浮点数 - date 日期类型  **约束限制** 不涉及 **取值范围**: - text - keyword - long - integer - double - float - date  **默认值** 不涉及
	Type *KeyIndexType `json:"type,omitempty"`
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
