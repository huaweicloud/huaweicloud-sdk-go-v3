package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 只支持string,integer,float
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
