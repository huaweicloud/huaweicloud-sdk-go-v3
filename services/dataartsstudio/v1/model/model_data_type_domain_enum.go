package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataTypeDomainEnum 字段类型所属域。 枚举值：   - NUMBER: 数字类型   - STRING: 字符类型   - DATETIME: 日期类型   - BLOB: 大对象(BLOB)   - OTHER: 其它类型
type DataTypeDomainEnum struct {
	value string
}

type DataTypeDomainEnumEnum struct {
	NUMBER   DataTypeDomainEnum
	STRING   DataTypeDomainEnum
	DATETIME DataTypeDomainEnum
	BLOB     DataTypeDomainEnum
	OTHER    DataTypeDomainEnum
}

func GetDataTypeDomainEnumEnum() DataTypeDomainEnumEnum {
	return DataTypeDomainEnumEnum{
		NUMBER: DataTypeDomainEnum{
			value: "NUMBER",
		},
		STRING: DataTypeDomainEnum{
			value: "STRING",
		},
		DATETIME: DataTypeDomainEnum{
			value: "DATETIME",
		},
		BLOB: DataTypeDomainEnum{
			value: "BLOB",
		},
		OTHER: DataTypeDomainEnum{
			value: "OTHER",
		},
	}
}

func (c DataTypeDomainEnum) Value() string {
	return c.value
}

func (c DataTypeDomainEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataTypeDomainEnum) UnmarshalJSON(b []byte) error {
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
