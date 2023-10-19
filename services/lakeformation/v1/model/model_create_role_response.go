package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRoleResponse Response Object
type CreateRoleResponse struct {

	// 角色名称。只能包含字母、数字和下划线，且长度为1~255个字符。
	RoleName *string `json:"role_name,omitempty"`

	// 描述信息。最大长度为4000个字符。当无描述信息时，则description值为null，当值为null时，响应Body无该参数。
	Description *string `json:"description,omitempty"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 AGENTTENANT 委托 OTHER 其它
	PrincipalSource *CreateRoleResponsePrincipalSource `json:"principal_source,omitempty"`
	HttpStatusCode  int                                `json:"-"`
}

func (o CreateRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoleResponse struct{}"
	}

	return strings.Join([]string{"CreateRoleResponse", string(data)}, " ")
}

type CreateRoleResponsePrincipalSource struct {
	value string
}

type CreateRoleResponsePrincipalSourceEnum struct {
	IAM         CreateRoleResponsePrincipalSource
	SAML        CreateRoleResponsePrincipalSource
	LDAP        CreateRoleResponsePrincipalSource
	LOCAL       CreateRoleResponsePrincipalSource
	AGENTTENANT CreateRoleResponsePrincipalSource
	OTHER       CreateRoleResponsePrincipalSource
}

func GetCreateRoleResponsePrincipalSourceEnum() CreateRoleResponsePrincipalSourceEnum {
	return CreateRoleResponsePrincipalSourceEnum{
		IAM: CreateRoleResponsePrincipalSource{
			value: "IAM",
		},
		SAML: CreateRoleResponsePrincipalSource{
			value: "SAML",
		},
		LDAP: CreateRoleResponsePrincipalSource{
			value: "LDAP",
		},
		LOCAL: CreateRoleResponsePrincipalSource{
			value: "LOCAL",
		},
		AGENTTENANT: CreateRoleResponsePrincipalSource{
			value: "AGENTTENANT",
		},
		OTHER: CreateRoleResponsePrincipalSource{
			value: "OTHER",
		},
	}
}

func (c CreateRoleResponsePrincipalSource) Value() string {
	return c.value
}

func (c CreateRoleResponsePrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRoleResponsePrincipalSource) UnmarshalJSON(b []byte) error {
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
