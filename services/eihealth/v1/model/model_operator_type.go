package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperatorType struct {
	value string
}

type OperatorTypeEnum struct {
	OR  OperatorType
	AND OperatorType
}

func GetOperatorTypeEnum() OperatorTypeEnum {
	return OperatorTypeEnum{
		OR: OperatorType{
			value: "or",
		},
		AND: OperatorType{
			value: "and",
		},
	}
}

func (c OperatorType) Value() string {
	return c.value
}

func (c OperatorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperatorType) UnmarshalJSON(b []byte) error {
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
