package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InputParaValueType 只支持string,integer,float
type InputParaValueType struct {
	value string
}

type InputParaValueTypeEnum struct {
	STRING  InputParaValueType
	INTEGER InputParaValueType
	FLOAT   InputParaValueType
}

func GetInputParaValueTypeEnum() InputParaValueTypeEnum {
	return InputParaValueTypeEnum{
		STRING: InputParaValueType{
			value: "string",
		},
		INTEGER: InputParaValueType{
			value: "integer",
		},
		FLOAT: InputParaValueType{
			value: "float",
		},
	}
}

func (c InputParaValueType) Value() string {
	return c.value
}

func (c InputParaValueType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InputParaValueType) UnmarshalJSON(b []byte) error {
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
