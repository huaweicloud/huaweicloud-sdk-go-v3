package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PlatformTypeEnum 平台类型： * `AD` - AD域 * `LOCAL` - LiteAs * `SYSTEM` - 系统内置
type PlatformTypeEnum struct {
	value string
}

type PlatformTypeEnumEnum struct {
	AD     PlatformTypeEnum
	LOCAL  PlatformTypeEnum
	SYSTEM PlatformTypeEnum
}

func GetPlatformTypeEnumEnum() PlatformTypeEnumEnum {
	return PlatformTypeEnumEnum{
		AD: PlatformTypeEnum{
			value: "AD",
		},
		LOCAL: PlatformTypeEnum{
			value: "LOCAL",
		},
		SYSTEM: PlatformTypeEnum{
			value: "SYSTEM",
		},
	}
}

func (c PlatformTypeEnum) Value() string {
	return c.value
}

func (c PlatformTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PlatformTypeEnum) UnmarshalJSON(b []byte) error {
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
