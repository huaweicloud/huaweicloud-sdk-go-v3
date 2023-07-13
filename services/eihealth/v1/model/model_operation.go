package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Operation struct {
	value string
}

type OperationEnum struct {
	OPEN  Operation
	CLOSE Operation
}

func GetOperationEnum() OperationEnum {
	return OperationEnum{
		OPEN: Operation{
			value: "OPEN",
		},
		CLOSE: Operation{
			value: "CLOSE",
		},
	}
}

func (c Operation) Value() string {
	return c.value
}

func (c Operation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Operation) UnmarshalJSON(b []byte) error {
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
