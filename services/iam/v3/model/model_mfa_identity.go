package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type MfaIdentity struct {
	// 认证方法，该字段内容为[\"password\", \"totp\"]。

	Methods []MfaIdentityMethods `json:"methods"`

	Password *PwdPassword `json:"password"`

	Totp *MfaTotp `json:"totp"`
}

func (o MfaIdentity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaIdentity struct{}"
	}

	return strings.Join([]string{"MfaIdentity", string(data)}, " ")
}

type MfaIdentityMethods struct {
	value string
}

type MfaIdentityMethodsEnum struct {
	PASSWORD MfaIdentityMethods
	TOTP     MfaIdentityMethods
}

func GetMfaIdentityMethodsEnum() MfaIdentityMethodsEnum {
	return MfaIdentityMethodsEnum{
		PASSWORD: MfaIdentityMethods{
			value: "password",
		},
		TOTP: MfaIdentityMethods{
			value: " totp",
		},
	}
}

func (c MfaIdentityMethods) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MfaIdentityMethods) UnmarshalJSON(b []byte) error {
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
