package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HostAuthorizationBody 登录主机鉴权，使用密码登录则填写密码即可，使用密钥则填写密钥，二选一即可。
type HostAuthorizationBody struct {

	// 用户名，可输入中英文，数字和符号(-_.)。
	Username string `json:"username"`

	// 密码，认证类型为0时，密码必填。
	Password *string `json:"password,omitempty"`

	// 密钥，认证类型为1时，密钥必填
	PrivateKey *string `json:"private_key,omitempty"`

	// 认证类型，0表示使用密码认证，1表示使用密钥认证
	TrustedType HostAuthorizationBodyTrustedType `json:"trusted_type"`
}

func (o HostAuthorizationBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostAuthorizationBody struct{}"
	}

	return strings.Join([]string{"HostAuthorizationBody", string(data)}, " ")
}

type HostAuthorizationBodyTrustedType struct {
	value int32
}

type HostAuthorizationBodyTrustedTypeEnum struct {
	E_0 HostAuthorizationBodyTrustedType
	E_1 HostAuthorizationBodyTrustedType
}

func GetHostAuthorizationBodyTrustedTypeEnum() HostAuthorizationBodyTrustedTypeEnum {
	return HostAuthorizationBodyTrustedTypeEnum{
		E_0: HostAuthorizationBodyTrustedType{
			value: 0,
		}, E_1: HostAuthorizationBodyTrustedType{
			value: 1,
		},
	}
}

func (c HostAuthorizationBodyTrustedType) Value() int32 {
	return c.value
}

func (c HostAuthorizationBodyTrustedType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostAuthorizationBodyTrustedType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
