package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Role 角色响应
type Role struct {

	// 角色名称。只能包含字母、数字和下划线，且长度为1~255个字符。
	RoleName string `json:"role_name"`

	// 描述信息。最大长度为4000个字符。当无描述信息时，则description值为null，当值为null时，响应Body无该参数。
	Description *string `json:"description,omitempty"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 AGENTTENANT 委托 OTHER 其它
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
	IAM         RolePrincipalSource
	SAML        RolePrincipalSource
	LDAP        RolePrincipalSource
	LOCAL       RolePrincipalSource
	AGENTTENANT RolePrincipalSource
	OTHER       RolePrincipalSource
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
		AGENTTENANT: RolePrincipalSource{
			value: "AGENTTENANT",
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
