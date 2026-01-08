package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthTypeEnum 认证类型。 RADIUS_GATEWAY：短信辅助认证 OAUTH2:OAUTH2认证 CLIENT_TOKEN:本地token认证 USER_PASSWORD:密码认证类型 SAML2:SAML 2.0 联邦认证
type AuthTypeEnum struct {
	value string
}

type AuthTypeEnumEnum struct {
	RADIUS_GATEWAY AuthTypeEnum
	OAUTH2         AuthTypeEnum
	LDAP           AuthTypeEnum
	CLIENT_TOKEN   AuthTypeEnum
	USER_PASSWORD  AuthTypeEnum
	FINGER         AuthTypeEnum
	SAML2          AuthTypeEnum
}

func GetAuthTypeEnumEnum() AuthTypeEnumEnum {
	return AuthTypeEnumEnum{
		RADIUS_GATEWAY: AuthTypeEnum{
			value: "RADIUS_GATEWAY",
		},
		OAUTH2: AuthTypeEnum{
			value: "OAUTH2",
		},
		LDAP: AuthTypeEnum{
			value: "LDAP",
		},
		CLIENT_TOKEN: AuthTypeEnum{
			value: "CLIENT_TOKEN",
		},
		USER_PASSWORD: AuthTypeEnum{
			value: "USER_PASSWORD",
		},
		FINGER: AuthTypeEnum{
			value: "FINGER",
		},
		SAML2: AuthTypeEnum{
			value: "SAML2",
		},
	}
}

func (c AuthTypeEnum) Value() string {
	return c.value
}

func (c AuthTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthTypeEnum) UnmarshalJSON(b []byte) error {
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
