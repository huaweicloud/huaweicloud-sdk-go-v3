package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthorizationTypeEnum 授权类型，基于应用(组)授权，默认为APP_GROUP授权。 * 'APP' - 按照应用授权 * 'APP_GROUP' - 按照应用组授权
type AuthorizationTypeEnum struct {
	value string
}

type AuthorizationTypeEnumEnum struct {
	APP       AuthorizationTypeEnum
	APP_GROUP AuthorizationTypeEnum
}

func GetAuthorizationTypeEnumEnum() AuthorizationTypeEnumEnum {
	return AuthorizationTypeEnumEnum{
		APP: AuthorizationTypeEnum{
			value: "APP",
		},
		APP_GROUP: AuthorizationTypeEnum{
			value: "APP_GROUP",
		},
	}
}

func (c AuthorizationTypeEnum) Value() string {
	return c.value
}

func (c AuthorizationTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizationTypeEnum) UnmarshalJSON(b []byte) error {
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
