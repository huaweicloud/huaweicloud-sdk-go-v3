package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CatalogInput catalog信息
type CatalogInput struct {

	// catalog名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	CatalogName string `json:"catalog_name"`

	// 描述信息。最大长度为4000个字符。
	Description *string `json:"description,omitempty"`

	// 路径地址。例如obs://location/uri/
	Location *string `json:"location,omitempty"`

	// 数据库路径列表。最小条目数为0，最大条目数为1000。
	DatabaseLocationList *[]string `json:"database_location_list,omitempty"`

	// 分支名称。只能包含字母、数字和下划线，且长度为1~32个字符。
	BranchName *string `json:"branch_name,omitempty"`

	// catalog所有者。只能包含字母、数字和下划线，且最大长度为128个字符。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色
	OwnerType *CatalogInputOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它
	OwnerSource *CatalogInputOwnerSource `json:"owner_source,omitempty"`
}

func (o CatalogInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogInput struct{}"
	}

	return strings.Join([]string{"CatalogInput", string(data)}, " ")
}

type CatalogInputOwnerType struct {
	value string
}

type CatalogInputOwnerTypeEnum struct {
	USER  CatalogInputOwnerType
	ROLE  CatalogInputOwnerType
	GROUP CatalogInputOwnerType
}

func GetCatalogInputOwnerTypeEnum() CatalogInputOwnerTypeEnum {
	return CatalogInputOwnerTypeEnum{
		USER: CatalogInputOwnerType{
			value: "USER",
		},
		ROLE: CatalogInputOwnerType{
			value: "ROLE",
		},
		GROUP: CatalogInputOwnerType{
			value: "GROUP",
		},
	}
}

func (c CatalogInputOwnerType) Value() string {
	return c.value
}

func (c CatalogInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CatalogInputOwnerType) UnmarshalJSON(b []byte) error {
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

type CatalogInputOwnerSource struct {
	value string
}

type CatalogInputOwnerSourceEnum struct {
	IAM         CatalogInputOwnerSource
	SAML        CatalogInputOwnerSource
	LDAP        CatalogInputOwnerSource
	LOCAL       CatalogInputOwnerSource
	AGENTTENANT CatalogInputOwnerSource
	OTHER       CatalogInputOwnerSource
}

func GetCatalogInputOwnerSourceEnum() CatalogInputOwnerSourceEnum {
	return CatalogInputOwnerSourceEnum{
		IAM: CatalogInputOwnerSource{
			value: "IAM",
		},
		SAML: CatalogInputOwnerSource{
			value: "SAML",
		},
		LDAP: CatalogInputOwnerSource{
			value: "LDAP",
		},
		LOCAL: CatalogInputOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: CatalogInputOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: CatalogInputOwnerSource{
			value: "OTHER",
		},
	}
}

func (c CatalogInputOwnerSource) Value() string {
	return c.value
}

func (c CatalogInputOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CatalogInputOwnerSource) UnmarshalJSON(b []byte) error {
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
