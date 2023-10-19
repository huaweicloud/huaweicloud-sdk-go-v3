package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRoleResponse Response Object
type UpdateRoleResponse struct {

	// 角色名称。只能包含字母、数字和下划线，且长度为1~255个字符。
	RoleName *string `json:"role_name,omitempty"`

	// 描述信息。最大长度为4000个字符。当无描述信息时，则description值为null，当值为null时，响应Body无该参数。
	Description *string `json:"description,omitempty"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 AGENTTENANT 委托 OTHER 其它
	PrincipalSource *UpdateRoleResponsePrincipalSource `json:"principal_source,omitempty"`
	HttpStatusCode  int                                `json:"-"`
}

func (o UpdateRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoleResponse struct{}"
	}

	return strings.Join([]string{"UpdateRoleResponse", string(data)}, " ")
}

type UpdateRoleResponsePrincipalSource struct {
	value string
}

type UpdateRoleResponsePrincipalSourceEnum struct {
	IAM         UpdateRoleResponsePrincipalSource
	SAML        UpdateRoleResponsePrincipalSource
	LDAP        UpdateRoleResponsePrincipalSource
	LOCAL       UpdateRoleResponsePrincipalSource
	AGENTTENANT UpdateRoleResponsePrincipalSource
	OTHER       UpdateRoleResponsePrincipalSource
}

func GetUpdateRoleResponsePrincipalSourceEnum() UpdateRoleResponsePrincipalSourceEnum {
	return UpdateRoleResponsePrincipalSourceEnum{
		IAM: UpdateRoleResponsePrincipalSource{
			value: "IAM",
		},
		SAML: UpdateRoleResponsePrincipalSource{
			value: "SAML",
		},
		LDAP: UpdateRoleResponsePrincipalSource{
			value: "LDAP",
		},
		LOCAL: UpdateRoleResponsePrincipalSource{
			value: "LOCAL",
		},
		AGENTTENANT: UpdateRoleResponsePrincipalSource{
			value: "AGENTTENANT",
		},
		OTHER: UpdateRoleResponsePrincipalSource{
			value: "OTHER",
		},
	}
}

func (c UpdateRoleResponsePrincipalSource) Value() string {
	return c.value
}

func (c UpdateRoleResponsePrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRoleResponsePrincipalSource) UnmarshalJSON(b []byte) error {
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
