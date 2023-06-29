package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UserRole 角色拥有的用户/用户模型
type UserRole struct {

	// 主体类型 USER用户 GROUP组
	PrincipalType UserRolePrincipalType `json:"principal_type"`

	// 主体名称
	PrincipalName string `json:"principal_name"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 OTHER 其它
	PrincipalSource UserRolePrincipalSource `json:"principal_source"`
}

func (o UserRole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserRole struct{}"
	}

	return strings.Join([]string{"UserRole", string(data)}, " ")
}

type UserRolePrincipalType struct {
	value string
}

type UserRolePrincipalTypeEnum struct {
	USER  UserRolePrincipalType
	GROUP UserRolePrincipalType
}

func GetUserRolePrincipalTypeEnum() UserRolePrincipalTypeEnum {
	return UserRolePrincipalTypeEnum{
		USER: UserRolePrincipalType{
			value: "USER",
		},
		GROUP: UserRolePrincipalType{
			value: "GROUP",
		},
	}
}

func (c UserRolePrincipalType) Value() string {
	return c.value
}

func (c UserRolePrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserRolePrincipalType) UnmarshalJSON(b []byte) error {
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

type UserRolePrincipalSource struct {
	value string
}

type UserRolePrincipalSourceEnum struct {
	IAM   UserRolePrincipalSource
	SAML  UserRolePrincipalSource
	LDAP  UserRolePrincipalSource
	LOCAL UserRolePrincipalSource
	OTHER UserRolePrincipalSource
}

func GetUserRolePrincipalSourceEnum() UserRolePrincipalSourceEnum {
	return UserRolePrincipalSourceEnum{
		IAM: UserRolePrincipalSource{
			value: "IAM",
		},
		SAML: UserRolePrincipalSource{
			value: "SAML",
		},
		LDAP: UserRolePrincipalSource{
			value: "LDAP",
		},
		LOCAL: UserRolePrincipalSource{
			value: "LOCAL",
		},
		OTHER: UserRolePrincipalSource{
			value: "OTHER",
		},
	}
}

func (c UserRolePrincipalSource) Value() string {
	return c.value
}

func (c UserRolePrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserRolePrincipalSource) UnmarshalJSON(b []byte) error {
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
