package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateRoleResponse struct {

	// role名字
	RoleName *string `json:"role_name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 主体来源 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 OTHER 其它
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
	IAM   UpdateRoleResponsePrincipalSource
	SAML  UpdateRoleResponsePrincipalSource
	LDAP  UpdateRoleResponsePrincipalSource
	LOCAL UpdateRoleResponsePrincipalSource
	OTHER UpdateRoleResponsePrincipalSource
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
