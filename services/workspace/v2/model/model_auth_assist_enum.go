package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 认证类型。 OTP：OTP辅助认证
type AuthAssistEnum struct {
	value string
}

type AuthAssistEnumEnum struct {
	OTP            AuthAssistEnum
	RADIUS_GATEWAY AuthAssistEnum
	RADIUS         AuthAssistEnum
}

func GetAuthAssistEnumEnum() AuthAssistEnumEnum {
	return AuthAssistEnumEnum{
		OTP: AuthAssistEnum{
			value: "OTP",
		},
		RADIUS_GATEWAY: AuthAssistEnum{
			value: "RADIUS_GATEWAY",
		},
		RADIUS: AuthAssistEnum{
			value: "RADIUS",
		},
	}
}

func (c AuthAssistEnum) Value() string {
	return c.value
}

func (c AuthAssistEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthAssistEnum) UnmarshalJSON(b []byte) error {
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
