package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StrongConstraintType struct {
	value string
}

type StrongConstraintTypeEnum struct {
	BOOL   StrongConstraintType
	RANGE  StrongConstraintType
	STRUCT StrongConstraintType
}

func GetStrongConstraintTypeEnum() StrongConstraintTypeEnum {
	return StrongConstraintTypeEnum{
		BOOL: StrongConstraintType{
			value: "bool",
		},
		RANGE: StrongConstraintType{
			value: "range",
		},
		STRUCT: StrongConstraintType{
			value: "struct",
		},
	}
}

func (c StrongConstraintType) Value() string {
	return c.value
}

func (c StrongConstraintType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StrongConstraintType) UnmarshalJSON(b []byte) error {
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
