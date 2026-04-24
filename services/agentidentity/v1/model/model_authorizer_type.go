package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthorizerType The authorizer type.
type AuthorizerType struct {
	value string
}

type AuthorizerTypeEnum struct {
	CUSTOM_JWT AuthorizerType
	IAM        AuthorizerType
	API_KEY    AuthorizerType
	NONE       AuthorizerType
}

func GetAuthorizerTypeEnum() AuthorizerTypeEnum {
	return AuthorizerTypeEnum{
		CUSTOM_JWT: AuthorizerType{
			value: "CUSTOM_JWT",
		},
		IAM: AuthorizerType{
			value: "IAM",
		},
		API_KEY: AuthorizerType{
			value: "API_KEY",
		},
		NONE: AuthorizerType{
			value: "NONE",
		},
	}
}

func (c AuthorizerType) Value() string {
	return c.value
}

func (c AuthorizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizerType) UnmarshalJSON(b []byte) error {
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
