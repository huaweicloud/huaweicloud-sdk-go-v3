package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConnectionTypeEnum 中心网络连接类型定义： - ER-ER (ER-ER Connection) - ER-GDGW (ER-GDGW Attachment Connection) - ER-ER_ROUTE_TABLE (ER-ER_ROUTE_TABLE Attachment Connection)
type ConnectionTypeEnum struct {
	value string
}

type ConnectionTypeEnumEnum struct {
	ER_ER             ConnectionTypeEnum
	ER_GDGW           ConnectionTypeEnum
	ER_ER_ROUTE_TABLE ConnectionTypeEnum
}

func GetConnectionTypeEnumEnum() ConnectionTypeEnumEnum {
	return ConnectionTypeEnumEnum{
		ER_ER: ConnectionTypeEnum{
			value: "ER-ER",
		},
		ER_GDGW: ConnectionTypeEnum{
			value: "ER-GDGW",
		},
		ER_ER_ROUTE_TABLE: ConnectionTypeEnum{
			value: "ER-ER_ROUTE_TABLE",
		},
	}
}

func (c ConnectionTypeEnum) Value() string {
	return c.value
}

func (c ConnectionTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionTypeEnum) UnmarshalJSON(b []byte) error {
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
