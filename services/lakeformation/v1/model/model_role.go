package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Role role 响应
type Role struct {

	// role名字
	RoleName string `json:"role_name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 OTHER 其它
	PrincipalSource RolePrincipalSource `json:"principal_source"`
}

func (o Role) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Role struct{}"
	}

	return strings.Join([]string{"Role", string(data)}, " ")
}

type RolePrincipalSource struct {
	value string
}

type RolePrincipalSourceEnum struct {
	IAM   RolePrincipalSource
	SAML  RolePrincipalSource
	LDAP  RolePrincipalSource
	LOCAL RolePrincipalSource
	OTHER RolePrincipalSource
}

func GetRolePrincipalSourceEnum() RolePrincipalSourceEnum {
	return RolePrincipalSourceEnum{
		IAM: RolePrincipalSource{
			value: "IAM",
		},
		SAML: RolePrincipalSource{
			value: "SAML",
		},
		LDAP: RolePrincipalSource{
			value: "LDAP",
		},
		LOCAL: RolePrincipalSource{
			value: "LOCAL",
		},
		OTHER: RolePrincipalSource{
			value: "OTHER",
		},
	}
}

func (c RolePrincipalSource) Value() string {
	return c.value
}

func (c RolePrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RolePrincipalSource) UnmarshalJSON(b []byte) error {
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
