package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Principal struct {

	// 主体类型 USER用户 GROUP组 ROLE角色 SHARE共享 OTHER其它
	PrincipalType PrincipalPrincipalType `json:"principal_type"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户
	PrincipalSource PrincipalPrincipalSource `json:"principal_source"`

	// 主体名
	PrincipalName string `json:"principal_name"`
}

func (o Principal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Principal struct{}"
	}

	return strings.Join([]string{"Principal", string(data)}, " ")
}

type PrincipalPrincipalType struct {
	value string
}

type PrincipalPrincipalTypeEnum struct {
	USER  PrincipalPrincipalType
	GROUP PrincipalPrincipalType
	ROLE  PrincipalPrincipalType
	SHARE PrincipalPrincipalType
	OTHER PrincipalPrincipalType
}

func GetPrincipalPrincipalTypeEnum() PrincipalPrincipalTypeEnum {
	return PrincipalPrincipalTypeEnum{
		USER: PrincipalPrincipalType{
			value: "USER",
		},
		GROUP: PrincipalPrincipalType{
			value: "GROUP",
		},
		ROLE: PrincipalPrincipalType{
			value: "ROLE",
		},
		SHARE: PrincipalPrincipalType{
			value: "SHARE",
		},
		OTHER: PrincipalPrincipalType{
			value: "OTHER",
		},
	}
}

func (c PrincipalPrincipalType) Value() string {
	return c.value
}

func (c PrincipalPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrincipalPrincipalType) UnmarshalJSON(b []byte) error {
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

type PrincipalPrincipalSource struct {
	value string
}

type PrincipalPrincipalSourceEnum struct {
	IAM   PrincipalPrincipalSource
	SAML  PrincipalPrincipalSource
	LDAP  PrincipalPrincipalSource
	LOCAL PrincipalPrincipalSource
}

func GetPrincipalPrincipalSourceEnum() PrincipalPrincipalSourceEnum {
	return PrincipalPrincipalSourceEnum{
		IAM: PrincipalPrincipalSource{
			value: "IAM",
		},
		SAML: PrincipalPrincipalSource{
			value: "SAML",
		},
		LDAP: PrincipalPrincipalSource{
			value: "LDAP",
		},
		LOCAL: PrincipalPrincipalSource{
			value: "LOCAL",
		},
	}
}

func (c PrincipalPrincipalSource) Value() string {
	return c.value
}

func (c PrincipalPrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrincipalPrincipalSource) UnmarshalJSON(b []byte) error {
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
