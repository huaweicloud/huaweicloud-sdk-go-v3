package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 登录主机鉴权，使用密码登录则填写密码即可，使用密钥则填写密钥，二选一即可。
type DeploymentHostAuthorizationBody struct {

	// 用户名，可输入中英文，数字和符号(-_.)。
	Username string `json:"username"`

	// 密码，认证类型为0时，密码必填。
	Password *string `json:"password,omitempty"`

	// 密钥，认证类型为1时，密钥必填
	PrivateKey *string `json:"private_key,omitempty"`

	// 认证类型，0表示使用密码认证，1表示使用密钥认证
	TrustedType DeploymentHostAuthorizationBodyTrustedType `json:"trusted_type"`
}

func (o DeploymentHostAuthorizationBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentHostAuthorizationBody struct{}"
	}

	return strings.Join([]string{"DeploymentHostAuthorizationBody", string(data)}, " ")
}

type DeploymentHostAuthorizationBodyTrustedType struct {
	value int32
}

type DeploymentHostAuthorizationBodyTrustedTypeEnum struct {
	E_0 DeploymentHostAuthorizationBodyTrustedType
	E_1 DeploymentHostAuthorizationBodyTrustedType
}

func GetDeploymentHostAuthorizationBodyTrustedTypeEnum() DeploymentHostAuthorizationBodyTrustedTypeEnum {
	return DeploymentHostAuthorizationBodyTrustedTypeEnum{
		E_0: DeploymentHostAuthorizationBodyTrustedType{
			value: 0,
		}, E_1: DeploymentHostAuthorizationBodyTrustedType{
			value: 1,
		},
	}
}

func (c DeploymentHostAuthorizationBodyTrustedType) Value() int32 {
	return c.value
}

func (c DeploymentHostAuthorizationBodyTrustedType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeploymentHostAuthorizationBodyTrustedType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
