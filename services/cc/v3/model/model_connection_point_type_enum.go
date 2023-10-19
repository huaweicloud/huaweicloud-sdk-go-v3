package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConnectionPointTypeEnum 中心网络连接点类型定义： - ER (EnterpriseRouter) - GDGW (Global DC Gateway) - ER_ROUTE_TABLE (Route Table)
type ConnectionPointTypeEnum struct {
	value string
}

type ConnectionPointTypeEnumEnum struct {
	ER             ConnectionPointTypeEnum
	GDGW           ConnectionPointTypeEnum
	ER_ROUTE_TABLE ConnectionPointTypeEnum
}

func GetConnectionPointTypeEnumEnum() ConnectionPointTypeEnumEnum {
	return ConnectionPointTypeEnumEnum{
		ER: ConnectionPointTypeEnum{
			value: "ER",
		},
		GDGW: ConnectionPointTypeEnum{
			value: "GDGW",
		},
		ER_ROUTE_TABLE: ConnectionPointTypeEnum{
			value: "ER_ROUTE_TABLE",
		},
	}
}

func (c ConnectionPointTypeEnum) Value() string {
	return c.value
}

func (c ConnectionPointTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionPointTypeEnum) UnmarshalJSON(b []byte) error {
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
