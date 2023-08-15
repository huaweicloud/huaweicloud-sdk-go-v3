package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OsTypeEnum 系统类型，当前仅支持Windows。 * `Linux` - * `Windows` - * `Other` -
type OsTypeEnum struct {
	value string
}

type OsTypeEnumEnum struct {
	LINUX   OsTypeEnum
	WINDOWS OsTypeEnum
	OTHER   OsTypeEnum
}

func GetOsTypeEnumEnum() OsTypeEnumEnum {
	return OsTypeEnumEnum{
		LINUX: OsTypeEnum{
			value: "Linux",
		},
		WINDOWS: OsTypeEnum{
			value: "Windows",
		},
		OTHER: OsTypeEnum{
			value: "Other",
		},
	}
}

func (c OsTypeEnum) Value() string {
	return c.value
}

func (c OsTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OsTypeEnum) UnmarshalJSON(b []byte) error {
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
