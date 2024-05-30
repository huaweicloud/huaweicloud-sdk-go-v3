package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EnvTypeEnum 开发生产环境类型。 枚举值：   - INVALID_TYPE: 无效环境   - DEV_TYPE: 开发环境   - PROD_TYPE: 生产环境   - DEV_PROD_TYPE: 开发生产环境
type EnvTypeEnum struct {
	value string
}

type EnvTypeEnumEnum struct {
	INVALID_TYPE  EnvTypeEnum
	DEV_TYPE      EnvTypeEnum
	PROD_TYPE     EnvTypeEnum
	DEV_PROD_TYPE EnvTypeEnum
}

func GetEnvTypeEnumEnum() EnvTypeEnumEnum {
	return EnvTypeEnumEnum{
		INVALID_TYPE: EnvTypeEnum{
			value: "INVALID_TYPE",
		},
		DEV_TYPE: EnvTypeEnum{
			value: "DEV_TYPE",
		},
		PROD_TYPE: EnvTypeEnum{
			value: "PROD_TYPE",
		},
		DEV_PROD_TYPE: EnvTypeEnum{
			value: "DEV_PROD_TYPE",
		},
	}
}

func (c EnvTypeEnum) Value() string {
	return c.value
}

func (c EnvTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvTypeEnum) UnmarshalJSON(b []byte) error {
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
