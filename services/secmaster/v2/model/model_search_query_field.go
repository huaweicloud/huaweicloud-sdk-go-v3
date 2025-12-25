package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SearchQueryField struct {

	// 字段名
	Name string `json:"name"`

	// **参数解释**: 数据类型 - boolean 布尔型 - byte 字节型 - short 短整型 - integer 整型 - long 长整型 - float 单精度浮点型 - half_float 半精度浮点型 - scaled_float 缩放浮点型 - double 双精度浮点型 - keyword 关键字型 - text 文本型 - date 日期型 - ip IP地址型 - binary 二进制型 - object 对象型 - nested 嵌套型  **约束限制** 不涉及 **取值范围**: - boolean - byte - short - integer - long - float - half_float - scaled_float - double - keyword - text - date - ip - binary - object - nested  **默认值** 不涉及
	Type SearchQueryFieldType `json:"type"`

	// 字段别名
	Alias *string `json:"alias,omitempty"`
}

func (o SearchQueryField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryField struct{}"
	}

	return strings.Join([]string{"SearchQueryField", string(data)}, " ")
}

type SearchQueryFieldType struct {
	value string
}

type SearchQueryFieldTypeEnum struct {
	BOOLEAN      SearchQueryFieldType
	BYTE         SearchQueryFieldType
	SHORT        SearchQueryFieldType
	INTEGER      SearchQueryFieldType
	LONG         SearchQueryFieldType
	FLOAT        SearchQueryFieldType
	HALF_FLOAT   SearchQueryFieldType
	SCALED_FLOAT SearchQueryFieldType
	DOUBLE       SearchQueryFieldType
	KEYWORD      SearchQueryFieldType
	TEXT         SearchQueryFieldType
	DATE         SearchQueryFieldType
	IP           SearchQueryFieldType
	BINARY       SearchQueryFieldType
	OBJECT       SearchQueryFieldType
	NESTED       SearchQueryFieldType
}

func GetSearchQueryFieldTypeEnum() SearchQueryFieldTypeEnum {
	return SearchQueryFieldTypeEnum{
		BOOLEAN: SearchQueryFieldType{
			value: "boolean",
		},
		BYTE: SearchQueryFieldType{
			value: "byte",
		},
		SHORT: SearchQueryFieldType{
			value: "short",
		},
		INTEGER: SearchQueryFieldType{
			value: "integer",
		},
		LONG: SearchQueryFieldType{
			value: "long",
		},
		FLOAT: SearchQueryFieldType{
			value: "float",
		},
		HALF_FLOAT: SearchQueryFieldType{
			value: "half_float",
		},
		SCALED_FLOAT: SearchQueryFieldType{
			value: "scaled_float",
		},
		DOUBLE: SearchQueryFieldType{
			value: "double",
		},
		KEYWORD: SearchQueryFieldType{
			value: "keyword",
		},
		TEXT: SearchQueryFieldType{
			value: "text",
		},
		DATE: SearchQueryFieldType{
			value: "date",
		},
		IP: SearchQueryFieldType{
			value: "ip",
		},
		BINARY: SearchQueryFieldType{
			value: "binary",
		},
		OBJECT: SearchQueryFieldType{
			value: "object",
		},
		NESTED: SearchQueryFieldType{
			value: "nested",
		},
	}
}

func (c SearchQueryFieldType) Value() string {
	return c.value
}

func (c SearchQueryFieldType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchQueryFieldType) UnmarshalJSON(b []byte) error {
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
