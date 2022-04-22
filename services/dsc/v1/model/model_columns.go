package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 字段类型列表，最大长度100
type Columns struct {

	// 数据的字段名称，最大支持长度256
	Name string `json:"name"`

	// 数据的字段类型
	Type ColumnsType `json:"type"`

	// 标记该字段是否为主键。true为主键，表示用来定位水印位置；false为非主键，将在该列嵌入/提取水印内容。字段类型列表中可同时包含多个为true或为false的字段
	PrimaryKey bool `json:"primary_key"`
}

func (o Columns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Columns struct{}"
	}

	return strings.Join([]string{"Columns", string(data)}, " ")
}

type ColumnsType struct {
	value string
}

type ColumnsTypeEnum struct {
	INTEGER ColumnsType
	STRING  ColumnsType
	DOUBLE  ColumnsType
}

func GetColumnsTypeEnum() ColumnsTypeEnum {
	return ColumnsTypeEnum{
		INTEGER: ColumnsType{
			value: "INTEGER",
		},
		STRING: ColumnsType{
			value: "STRING",
		},
		DOUBLE: ColumnsType{
			value: "DOUBLE",
		},
	}
}

func (c ColumnsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ColumnsType) UnmarshalJSON(b []byte) error {
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
