package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCatalogResponse Response Object
type CreateCatalogResponse struct {

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
	OwnerType *CreateCatalogResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerSource    *CreateCatalogResponseOwnerSource `json:"owner_source,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o CreateCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCatalogResponse struct{}"
	}

	return strings.Join([]string{"CreateCatalogResponse", string(data)}, " ")
}

type CreateCatalogResponseOwnerType struct {
	value string
}

type CreateCatalogResponseOwnerTypeEnum struct {
	USER  CreateCatalogResponseOwnerType
	ROLE  CreateCatalogResponseOwnerType
	GROUP CreateCatalogResponseOwnerType
}

func GetCreateCatalogResponseOwnerTypeEnum() CreateCatalogResponseOwnerTypeEnum {
	return CreateCatalogResponseOwnerTypeEnum{
		USER: CreateCatalogResponseOwnerType{
			value: "USER",
		},
		ROLE: CreateCatalogResponseOwnerType{
			value: "ROLE",
		},
		GROUP: CreateCatalogResponseOwnerType{
			value: "GROUP",
		},
	}
}

func (c CreateCatalogResponseOwnerType) Value() string {
	return c.value
}

func (c CreateCatalogResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCatalogResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type CreateCatalogResponseOwnerSource struct {
	value string
}

type CreateCatalogResponseOwnerSourceEnum struct {
	IAM         CreateCatalogResponseOwnerSource
	SAML        CreateCatalogResponseOwnerSource
	LDAP        CreateCatalogResponseOwnerSource
	LOCAL       CreateCatalogResponseOwnerSource
	AGENTTENANT CreateCatalogResponseOwnerSource
	OTHER       CreateCatalogResponseOwnerSource
}

func GetCreateCatalogResponseOwnerSourceEnum() CreateCatalogResponseOwnerSourceEnum {
	return CreateCatalogResponseOwnerSourceEnum{
		IAM: CreateCatalogResponseOwnerSource{
			value: "IAM",
		},
		SAML: CreateCatalogResponseOwnerSource{
			value: "SAML",
		},
		LDAP: CreateCatalogResponseOwnerSource{
			value: "LDAP",
		},
		LOCAL: CreateCatalogResponseOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: CreateCatalogResponseOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: CreateCatalogResponseOwnerSource{
			value: "OTHER",
		},
	}
}

func (c CreateCatalogResponseOwnerSource) Value() string {
	return c.value
}

func (c CreateCatalogResponseOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCatalogResponseOwnerSource) UnmarshalJSON(b []byte) error {
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
