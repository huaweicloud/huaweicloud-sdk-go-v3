package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RoleInfoInput 角色输入
type RoleInfoInput struct {

	// 角色名称。只能包含字母、数字和下划线，且长度为1~255个字符。
	RoleName string `json:"role_name"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 AGENTTENANT 委托 OTHER 其它
	PrincipalSource RoleInfoInputPrincipalSource `json:"principal_source"`
}

func (o RoleInfoInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleInfoInput struct{}"
	}

	return strings.Join([]string{"RoleInfoInput", string(data)}, " ")
}

type RoleInfoInputPrincipalSource struct {
	value string
}

type RoleInfoInputPrincipalSourceEnum struct {
	IAM         RoleInfoInputPrincipalSource
	SAML        RoleInfoInputPrincipalSource
	LDAP        RoleInfoInputPrincipalSource
	LOCAL       RoleInfoInputPrincipalSource
	AGENTTENANT RoleInfoInputPrincipalSource
	OTHER       RoleInfoInputPrincipalSource
}

func GetRoleInfoInputPrincipalSourceEnum() RoleInfoInputPrincipalSourceEnum {
	return RoleInfoInputPrincipalSourceEnum{
		IAM: RoleInfoInputPrincipalSource{
			value: "IAM",
		},
		SAML: RoleInfoInputPrincipalSource{
			value: "SAML",
		},
		LDAP: RoleInfoInputPrincipalSource{
			value: "LDAP",
		},
		LOCAL: RoleInfoInputPrincipalSource{
			value: "LOCAL",
		},
		AGENTTENANT: RoleInfoInputPrincipalSource{
			value: "AGENTTENANT",
		},
		OTHER: RoleInfoInputPrincipalSource{
			value: "OTHER",
		},
	}
}

func (c RoleInfoInputPrincipalSource) Value() string {
	return c.value
}

func (c RoleInfoInputPrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RoleInfoInputPrincipalSource) UnmarshalJSON(b []byte) error {
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
