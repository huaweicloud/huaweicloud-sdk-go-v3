package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 认证方式参数
type AuthOpt struct {
	// AppCode简易认证类型，仅在auth_type为APP时生效，默认为DISABLE： - DISABLE：不开启简易认证 - HEADER：开启简易认证且AppCode位置在HEADER

	AppCodeAuthType *AuthOptAppCodeAuthType `json:"app_code_auth_type,omitempty"`
}

func (o AuthOpt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthOpt struct{}"
	}

	return strings.Join([]string{"AuthOpt", string(data)}, " ")
}

type AuthOptAppCodeAuthType struct {
	value string
}

type AuthOptAppCodeAuthTypeEnum struct {
	DISABLE AuthOptAppCodeAuthType
	HEADER  AuthOptAppCodeAuthType
}

func GetAuthOptAppCodeAuthTypeEnum() AuthOptAppCodeAuthTypeEnum {
	return AuthOptAppCodeAuthTypeEnum{
		DISABLE: AuthOptAppCodeAuthType{
			value: "DISABLE",
		},
		HEADER: AuthOptAppCodeAuthType{
			value: "HEADER",
		},
	}
}

func (c AuthOptAppCodeAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthOptAppCodeAuthType) UnmarshalJSON(b []byte) error {
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
