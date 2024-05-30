package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SecretTypeEnum 密级类型。 枚举值：   - PUBLIC: 公开   - SECRET: 秘密   - CONFIDENTIAL: 机密   - SUPER_SECRET: 绝密
type SecretTypeEnum struct {
	value string
}

type SecretTypeEnumEnum struct {
	PUBLIC       SecretTypeEnum
	SECRET       SecretTypeEnum
	CONFIDENTIAL SecretTypeEnum
	SUPER_SECRET SecretTypeEnum
}

func GetSecretTypeEnumEnum() SecretTypeEnumEnum {
	return SecretTypeEnumEnum{
		PUBLIC: SecretTypeEnum{
			value: "PUBLIC",
		},
		SECRET: SecretTypeEnum{
			value: "SECRET",
		},
		CONFIDENTIAL: SecretTypeEnum{
			value: "CONFIDENTIAL",
		},
		SUPER_SECRET: SecretTypeEnum{
			value: "SUPER_SECRET",
		},
	}
}

func (c SecretTypeEnum) Value() string {
	return c.value
}

func (c SecretTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecretTypeEnum) UnmarshalJSON(b []byte) error {
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
