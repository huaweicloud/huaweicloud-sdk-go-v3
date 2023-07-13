package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConnectionType 目标连接类型。目前支持webhook：http连接；kafka：华为云官方kafka实例
type ConnectionType struct {
	value string
}

type ConnectionTypeEnum struct {
	WEBHOOK ConnectionType
	KAFKA   ConnectionType
}

func GetConnectionTypeEnum() ConnectionTypeEnum {
	return ConnectionTypeEnum{
		WEBHOOK: ConnectionType{
			value: "WEBHOOK",
		},
		KAFKA: ConnectionType{
			value: "KAFKA",
		},
	}
}

func (c ConnectionType) Value() string {
	return c.value
}

func (c ConnectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionType) UnmarshalJSON(b []byte) error {
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
