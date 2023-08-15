package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Auth 可选字段，鉴权认证类型。替换时需要该参数，重置时不需要该参数。
type Auth struct {

	// 取值为枚举类型。
	Type *AuthType `json:"type,omitempty"`

	// - type为枚举值password时，key表示密码； - type为枚举值keypair时，key表示私钥；
	Key *string `json:"key,omitempty"`
}

func (o Auth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Auth struct{}"
	}

	return strings.Join([]string{"Auth", string(data)}, " ")
}

type AuthType struct {
	value string
}

type AuthTypeEnum struct {
	PASSWORD AuthType
	KEYPAIR  AuthType
}

func GetAuthTypeEnum() AuthTypeEnum {
	return AuthTypeEnum{
		PASSWORD: AuthType{
			value: "password",
		},
		KEYPAIR: AuthType{
			value: "keypair",
		},
	}
}

func (c AuthType) Value() string {
	return c.value
}

func (c AuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthType) UnmarshalJSON(b []byte) error {
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
