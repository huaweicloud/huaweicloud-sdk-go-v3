package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RowKey struct {

	// 通道内JSON数据的JSON属性名，用于生成HBase数据的rowkey。
	Value string `json:"value"`

	// 通道内JSON数据的JSON属性的类型名称。
	Type RowKeyType `json:"type"`
}

func (o RowKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RowKey struct{}"
	}

	return strings.Join([]string{"RowKey", string(data)}, " ")
}

type RowKeyType struct {
	value string
}

type RowKeyTypeEnum struct {
	BIGINT    RowKeyType
	DOUBLE    RowKeyType
	BOOLEAN   RowKeyType
	TIMESTAMP RowKeyType
	STRING    RowKeyType
	DECIMAL   RowKeyType
}

func GetRowKeyTypeEnum() RowKeyTypeEnum {
	return RowKeyTypeEnum{
		BIGINT: RowKeyType{
			value: "Bigint",
		},
		DOUBLE: RowKeyType{
			value: "Double",
		},
		BOOLEAN: RowKeyType{
			value: "Boolean",
		},
		TIMESTAMP: RowKeyType{
			value: "Timestamp",
		},
		STRING: RowKeyType{
			value: "String",
		},
		DECIMAL: RowKeyType{
			value: "Decimal",
		},
	}
}

func (c RowKeyType) Value() string {
	return c.value
}

func (c RowKeyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RowKeyType) UnmarshalJSON(b []byte) error {
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
