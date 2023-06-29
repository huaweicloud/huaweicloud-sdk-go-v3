package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// YesNoEnum 是、否 通用枚举
type YesNoEnum struct {
	value string
}

type YesNoEnumEnum struct {
	Y YesNoEnum
	N YesNoEnum
}

func GetYesNoEnumEnum() YesNoEnumEnum {
	return YesNoEnumEnum{
		Y: YesNoEnum{
			value: "Y",
		},
		N: YesNoEnum{
			value: "N",
		},
	}
}

func (c YesNoEnum) Value() string {
	return c.value
}

func (c YesNoEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *YesNoEnum) UnmarshalJSON(b []byte) error {
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
