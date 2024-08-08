package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WeakConstraintType struct {
	value string
}

type WeakConstraintTypeEnum struct {
	BOOL        WeakConstraintType
	RANGE       WeakConstraintType
	STRUCT      WeakConstraintType
	MINIMIZE    WeakConstraintType
	MAXIMIZE    WeakConstraintType
	INTERACTION WeakConstraintType
}

func GetWeakConstraintTypeEnum() WeakConstraintTypeEnum {
	return WeakConstraintTypeEnum{
		BOOL: WeakConstraintType{
			value: "bool",
		},
		RANGE: WeakConstraintType{
			value: "range",
		},
		STRUCT: WeakConstraintType{
			value: "struct",
		},
		MINIMIZE: WeakConstraintType{
			value: "minimize",
		},
		MAXIMIZE: WeakConstraintType{
			value: "maximize",
		},
		INTERACTION: WeakConstraintType{
			value: "interaction",
		},
	}
}

func (c WeakConstraintType) Value() string {
	return c.value
}

func (c WeakConstraintType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WeakConstraintType) UnmarshalJSON(b []byte) error {
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
