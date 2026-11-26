package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowCatalogResponse Response Object
type ShowCatalogResponse struct {

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
	OwnerType *ShowCatalogResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerSource *ShowCatalogResponseOwnerSource `json:"owner_source,omitempty"`

	// catalogID。
	CatalogId *string `json:"catalog_id,omitempty"`

	// catalog元数据创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// catalog类型，可选值为DEFAULT（代表默认的元数据类型）、CLICKHOUSE（代表CH类型），不设置时默认为DEFAULT。 创建时指定，不可修改。
	Type *ShowCatalogResponseType `json:"type,omitempty"`

	// catalog元数据最后一次修改时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 版本ID
	VersionId      *string `json:"version_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCatalogResponse struct{}"
	}

	return strings.Join([]string{"ShowCatalogResponse", string(data)}, " ")
}

type ShowCatalogResponseOwnerType struct {
	value string
}

type ShowCatalogResponseOwnerTypeEnum struct {
	USER  ShowCatalogResponseOwnerType
	ROLE  ShowCatalogResponseOwnerType
	GROUP ShowCatalogResponseOwnerType
}

func GetShowCatalogResponseOwnerTypeEnum() ShowCatalogResponseOwnerTypeEnum {
	return ShowCatalogResponseOwnerTypeEnum{
		USER: ShowCatalogResponseOwnerType{
			value: "USER",
		},
		ROLE: ShowCatalogResponseOwnerType{
			value: "ROLE",
		},
		GROUP: ShowCatalogResponseOwnerType{
			value: "GROUP",
		},
	}
}

func (c ShowCatalogResponseOwnerType) Value() string {
	return c.value
}

func (c ShowCatalogResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCatalogResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type ShowCatalogResponseOwnerSource struct {
	value string
}

type ShowCatalogResponseOwnerSourceEnum struct {
	IAM         ShowCatalogResponseOwnerSource
	SAML        ShowCatalogResponseOwnerSource
	LDAP        ShowCatalogResponseOwnerSource
	LOCAL       ShowCatalogResponseOwnerSource
	AGENTTENANT ShowCatalogResponseOwnerSource
	OTHER       ShowCatalogResponseOwnerSource
}

func GetShowCatalogResponseOwnerSourceEnum() ShowCatalogResponseOwnerSourceEnum {
	return ShowCatalogResponseOwnerSourceEnum{
		IAM: ShowCatalogResponseOwnerSource{
			value: "IAM",
		},
		SAML: ShowCatalogResponseOwnerSource{
			value: "SAML",
		},
		LDAP: ShowCatalogResponseOwnerSource{
			value: "LDAP",
		},
		LOCAL: ShowCatalogResponseOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: ShowCatalogResponseOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: ShowCatalogResponseOwnerSource{
			value: "OTHER",
		},
	}
}

func (c ShowCatalogResponseOwnerSource) Value() string {
	return c.value
}

func (c ShowCatalogResponseOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCatalogResponseOwnerSource) UnmarshalJSON(b []byte) error {
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

type ShowCatalogResponseType struct {
	value string
}

type ShowCatalogResponseTypeEnum struct {
	DEFAULT    ShowCatalogResponseType
	CLICKHOUSE ShowCatalogResponseType
}

func GetShowCatalogResponseTypeEnum() ShowCatalogResponseTypeEnum {
	return ShowCatalogResponseTypeEnum{
		DEFAULT: ShowCatalogResponseType{
			value: "DEFAULT",
		},
		CLICKHOUSE: ShowCatalogResponseType{
			value: "CLICKHOUSE",
		},
	}
}

func (c ShowCatalogResponseType) Value() string {
	return c.value
}

func (c ShowCatalogResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCatalogResponseType) UnmarshalJSON(b []byte) error {
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
