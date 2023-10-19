package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateCatalogResponse Response Object
type UpdateCatalogResponse struct {

	// catalog名称
	CatalogName *string `json:"catalog_name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 路径地址。例如obs://location/uri/
	Location *string `json:"location,omitempty"`

	// 数据库路径列表。当值为null时，响应Body无该参数。
	DatabaseLocationList *[]string `json:"database_location_list,omitempty"`

	// catalog所有者。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerType *UpdateCatalogResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerSource    *UpdateCatalogResponseOwnerSource `json:"owner_source,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o UpdateCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCatalogResponse struct{}"
	}

	return strings.Join([]string{"UpdateCatalogResponse", string(data)}, " ")
}

type UpdateCatalogResponseOwnerType struct {
	value string
}

type UpdateCatalogResponseOwnerTypeEnum struct {
	USER  UpdateCatalogResponseOwnerType
	ROLE  UpdateCatalogResponseOwnerType
	GROUP UpdateCatalogResponseOwnerType
}

func GetUpdateCatalogResponseOwnerTypeEnum() UpdateCatalogResponseOwnerTypeEnum {
	return UpdateCatalogResponseOwnerTypeEnum{
		USER: UpdateCatalogResponseOwnerType{
			value: "USER",
		},
		ROLE: UpdateCatalogResponseOwnerType{
			value: "ROLE",
		},
		GROUP: UpdateCatalogResponseOwnerType{
			value: "GROUP",
		},
	}
}

func (c UpdateCatalogResponseOwnerType) Value() string {
	return c.value
}

func (c UpdateCatalogResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCatalogResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type UpdateCatalogResponseOwnerSource struct {
	value string
}

type UpdateCatalogResponseOwnerSourceEnum struct {
	IAM         UpdateCatalogResponseOwnerSource
	SAML        UpdateCatalogResponseOwnerSource
	LDAP        UpdateCatalogResponseOwnerSource
	LOCAL       UpdateCatalogResponseOwnerSource
	AGENTTENANT UpdateCatalogResponseOwnerSource
	OTHER       UpdateCatalogResponseOwnerSource
}

func GetUpdateCatalogResponseOwnerSourceEnum() UpdateCatalogResponseOwnerSourceEnum {
	return UpdateCatalogResponseOwnerSourceEnum{
		IAM: UpdateCatalogResponseOwnerSource{
			value: "IAM",
		},
		SAML: UpdateCatalogResponseOwnerSource{
			value: "SAML",
		},
		LDAP: UpdateCatalogResponseOwnerSource{
			value: "LDAP",
		},
		LOCAL: UpdateCatalogResponseOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: UpdateCatalogResponseOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: UpdateCatalogResponseOwnerSource{
			value: "OTHER",
		},
	}
}

func (c UpdateCatalogResponseOwnerSource) Value() string {
	return c.value
}

func (c UpdateCatalogResponseOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCatalogResponseOwnerSource) UnmarshalJSON(b []byte) error {
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
