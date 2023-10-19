package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRoleResponse Response Object
type ShowRoleResponse struct {

	// 角色名称。只能包含字母、数字和下划线，且长度为1~255个字符。
	RoleName *string `json:"role_name,omitempty"`

	// 描述信息。最大长度为4000个字符。当无描述信息时，则description值为null，当值为null时，响应Body无该参数。
	Description *string `json:"description,omitempty"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 AGENTTENANT 委托 OTHER 其它
	PrincipalSource *ShowRoleResponsePrincipalSource `json:"principal_source,omitempty"`
	HttpStatusCode  int                              `json:"-"`
}

func (o ShowRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRoleResponse struct{}"
	}

	return strings.Join([]string{"ShowRoleResponse", string(data)}, " ")
}

type ShowRoleResponsePrincipalSource struct {
	value string
}

type ShowRoleResponsePrincipalSourceEnum struct {
	IAM         ShowRoleResponsePrincipalSource
	SAML        ShowRoleResponsePrincipalSource
	LDAP        ShowRoleResponsePrincipalSource
	LOCAL       ShowRoleResponsePrincipalSource
	AGENTTENANT ShowRoleResponsePrincipalSource
	OTHER       ShowRoleResponsePrincipalSource
}

func GetShowRoleResponsePrincipalSourceEnum() ShowRoleResponsePrincipalSourceEnum {
	return ShowRoleResponsePrincipalSourceEnum{
		IAM: ShowRoleResponsePrincipalSource{
			value: "IAM",
		},
		SAML: ShowRoleResponsePrincipalSource{
			value: "SAML",
		},
		LDAP: ShowRoleResponsePrincipalSource{
			value: "LDAP",
		},
		LOCAL: ShowRoleResponsePrincipalSource{
			value: "LOCAL",
		},
		AGENTTENANT: ShowRoleResponsePrincipalSource{
			value: "AGENTTENANT",
		},
		OTHER: ShowRoleResponsePrincipalSource{
			value: "OTHER",
		},
	}
}

func (c ShowRoleResponsePrincipalSource) Value() string {
	return c.value
}

func (c ShowRoleResponsePrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRoleResponsePrincipalSource) UnmarshalJSON(b []byte) error {
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
