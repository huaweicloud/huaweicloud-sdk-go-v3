package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Column struct {

	// 存储该通道数据的HBase表数据的列族名称。
	ColumnFamilyName string `json:"column_family_name"`

	// 存储该通道数据的HBase表数据的列名称。  取值范围：1～32，只能包含英文字母、数字和下划线。
	ColumnName string `json:"column_name"`

	// 通道内JSON数据的JSON属性名，用于生成HBase数据的列值。
	Value string `json:"value"`

	// 通道内JSON数据的JSON属性的类型名称。
	Type ColumnType `json:"type"`
}

func (o Column) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Column struct{}"
	}

	return strings.Join([]string{"Column", string(data)}, " ")
}

type ColumnType struct {
	value string
}

type ColumnTypeEnum struct {
	BIGINT    ColumnType
	DOUBLE    ColumnType
	BOOLEAN   ColumnType
	TIMESTAMP ColumnType
	STRING    ColumnType
	DECIMAL   ColumnType
}

func GetColumnTypeEnum() ColumnTypeEnum {
	return ColumnTypeEnum{
		BIGINT: ColumnType{
			value: "Bigint",
		},
		DOUBLE: ColumnType{
			value: "Double",
		},
		BOOLEAN: ColumnType{
			value: "Boolean",
		},
		TIMESTAMP: ColumnType{
			value: "Timestamp",
		},
		STRING: ColumnType{
			value: "String",
		},
		DECIMAL: ColumnType{
			value: "Decimal",
		},
	}
}

func (c ColumnType) Value() string {
	return c.value
}

func (c ColumnType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ColumnType) UnmarshalJSON(b []byte) error {
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
