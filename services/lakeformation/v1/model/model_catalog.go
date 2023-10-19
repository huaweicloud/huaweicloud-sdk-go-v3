package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Catalog catalog response
type Catalog struct {

	// catalog名称
	CatalogName string `json:"catalog_name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 路径地址。例如obs://location/uri/
	Location *string `json:"location,omitempty"`

	// 数据库路径列表。当值为null时，响应Body无该参数。
	DatabaseLocationList *[]string `json:"database_location_list,omitempty"`

	// catalog所有者。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerType *CatalogOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerSource *CatalogOwnerSource `json:"owner_source,omitempty"`
}

func (o Catalog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Catalog struct{}"
	}

	return strings.Join([]string{"Catalog", string(data)}, " ")
}

type CatalogOwnerType struct {
	value string
}

type CatalogOwnerTypeEnum struct {
	USER  CatalogOwnerType
	ROLE  CatalogOwnerType
	GROUP CatalogOwnerType
}

func GetCatalogOwnerTypeEnum() CatalogOwnerTypeEnum {
	return CatalogOwnerTypeEnum{
		USER: CatalogOwnerType{
			value: "USER",
		},
		ROLE: CatalogOwnerType{
			value: "ROLE",
		},
		GROUP: CatalogOwnerType{
			value: "GROUP",
		},
	}
}

func (c CatalogOwnerType) Value() string {
	return c.value
}

func (c CatalogOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CatalogOwnerType) UnmarshalJSON(b []byte) error {
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

type CatalogOwnerSource struct {
	value string
}

type CatalogOwnerSourceEnum struct {
	IAM         CatalogOwnerSource
	SAML        CatalogOwnerSource
	LDAP        CatalogOwnerSource
	LOCAL       CatalogOwnerSource
	AGENTTENANT CatalogOwnerSource
	OTHER       CatalogOwnerSource
}

func GetCatalogOwnerSourceEnum() CatalogOwnerSourceEnum {
	return CatalogOwnerSourceEnum{
		IAM: CatalogOwnerSource{
			value: "IAM",
		},
		SAML: CatalogOwnerSource{
			value: "SAML",
		},
		LDAP: CatalogOwnerSource{
			value: "LDAP",
		},
		LOCAL: CatalogOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: CatalogOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: CatalogOwnerSource{
			value: "OTHER",
		},
	}
}

func (c CatalogOwnerSource) Value() string {
	return c.value
}

func (c CatalogOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CatalogOwnerSource) UnmarshalJSON(b []byte) error {
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
