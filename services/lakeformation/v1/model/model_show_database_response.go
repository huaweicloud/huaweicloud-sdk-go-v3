package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowDatabaseResponse Response Object
type ShowDatabaseResponse struct {

	// catalog名称
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色
	OwnerType *ShowDatabaseResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者授权来源类型,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerAuthSourceType *ShowDatabaseResponseOwnerAuthSourceType `json:"owner_auth_source_type,omitempty"`

	// 数据库描述信息
	Description *string `json:"description,omitempty"`

	// 数据库路径地址。例如obs://location/uri/
	Location *string `json:"location,omitempty"`

	// 参数信息
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表路径列表。LakeFormation服务分为一期和二期，一期响应Body无该参数，二期默认为null。当值为null时，响应Body无该参数。
	TableLocationList *[]string `json:"table_location_list,omitempty"`

	// 函数路径列表。默认为null，当值为null时，响应Body无该参数。
	FunctionLocationList *[]string `json:"function_location_list,omitempty"`

	// 权限策略来源。可选值为：lakeformation、ranger、hive。
	AccessPolicySource *ShowDatabaseResponseAccessPolicySource `json:"access_policy_source,omitempty"`

	// catalogID。
	CatalogId *string `json:"catalog_id,omitempty"`

	// 连接名称。
	ConnectionName *string `json:"connection_name,omitempty"`

	// 数据概况统计开关。默认状态为关，若要开启database下table的数据概况统计，需要先打开database的数据概况开关。database此开关开启后，database下所有table都会默认开启此开关
	DataStatisticEnable *bool `json:"data_statistic_enable,omitempty"`

	// 数据库ID。
	DatabaseId *string `json:"database_id,omitempty"`

	// 数据库状态：0表示正常，1表示删除中，2表示已删除，3表示恢复中，4表示清除中。
	DatabaseStatus *int32 `json:"database_status,omitempty"`

	// 数据库类型：MANAGED_DATABASE内置数据库、FOREIGN_DATABASE外置数据库。
	DatabaseType *ShowDatabaseResponseDatabaseType `json:"database_type,omitempty"`

	// 用户端数据库id，创建时指定，不可修改。
	ExternalDatabaseId *string `json:"external_database_id,omitempty"`

	// 权限连接名称。绑定后由该权限连接代理鉴权。为空时保留原权限连接。
	PermissionConnectionName *string `json:"permission_connection_name,omitempty"`

	// 数据库元数据最后一次修改时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 版本ID
	VersionId      *string `json:"version_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseResponse", string(data)}, " ")
}

type ShowDatabaseResponseOwnerType struct {
	value string
}

type ShowDatabaseResponseOwnerTypeEnum struct {
	USER  ShowDatabaseResponseOwnerType
	GROUP ShowDatabaseResponseOwnerType
	ROLE  ShowDatabaseResponseOwnerType
}

func GetShowDatabaseResponseOwnerTypeEnum() ShowDatabaseResponseOwnerTypeEnum {
	return ShowDatabaseResponseOwnerTypeEnum{
		USER: ShowDatabaseResponseOwnerType{
			value: "USER",
		},
		GROUP: ShowDatabaseResponseOwnerType{
			value: "GROUP",
		},
		ROLE: ShowDatabaseResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c ShowDatabaseResponseOwnerType) Value() string {
	return c.value
}

func (c ShowDatabaseResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatabaseResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type ShowDatabaseResponseOwnerAuthSourceType struct {
	value string
}

type ShowDatabaseResponseOwnerAuthSourceTypeEnum struct {
	IAM         ShowDatabaseResponseOwnerAuthSourceType
	SAML        ShowDatabaseResponseOwnerAuthSourceType
	LDAP        ShowDatabaseResponseOwnerAuthSourceType
	LOCAL       ShowDatabaseResponseOwnerAuthSourceType
	AGENTTENANT ShowDatabaseResponseOwnerAuthSourceType
	OTHER       ShowDatabaseResponseOwnerAuthSourceType
}

func GetShowDatabaseResponseOwnerAuthSourceTypeEnum() ShowDatabaseResponseOwnerAuthSourceTypeEnum {
	return ShowDatabaseResponseOwnerAuthSourceTypeEnum{
		IAM: ShowDatabaseResponseOwnerAuthSourceType{
			value: "IAM",
		},
		SAML: ShowDatabaseResponseOwnerAuthSourceType{
			value: "SAML",
		},
		LDAP: ShowDatabaseResponseOwnerAuthSourceType{
			value: "LDAP",
		},
		LOCAL: ShowDatabaseResponseOwnerAuthSourceType{
			value: "LOCAL",
		},
		AGENTTENANT: ShowDatabaseResponseOwnerAuthSourceType{
			value: "AGENTTENANT",
		},
		OTHER: ShowDatabaseResponseOwnerAuthSourceType{
			value: "OTHER",
		},
	}
}

func (c ShowDatabaseResponseOwnerAuthSourceType) Value() string {
	return c.value
}

func (c ShowDatabaseResponseOwnerAuthSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatabaseResponseOwnerAuthSourceType) UnmarshalJSON(b []byte) error {
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

type ShowDatabaseResponseAccessPolicySource struct {
	value string
}

type ShowDatabaseResponseAccessPolicySourceEnum struct {
	LAKEFORMATION ShowDatabaseResponseAccessPolicySource
	RANGER        ShowDatabaseResponseAccessPolicySource
	HIVE          ShowDatabaseResponseAccessPolicySource
}

func GetShowDatabaseResponseAccessPolicySourceEnum() ShowDatabaseResponseAccessPolicySourceEnum {
	return ShowDatabaseResponseAccessPolicySourceEnum{
		LAKEFORMATION: ShowDatabaseResponseAccessPolicySource{
			value: "lakeformation",
		},
		RANGER: ShowDatabaseResponseAccessPolicySource{
			value: "ranger",
		},
		HIVE: ShowDatabaseResponseAccessPolicySource{
			value: "hive",
		},
	}
}

func (c ShowDatabaseResponseAccessPolicySource) Value() string {
	return c.value
}

func (c ShowDatabaseResponseAccessPolicySource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatabaseResponseAccessPolicySource) UnmarshalJSON(b []byte) error {
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

type ShowDatabaseResponseDatabaseType struct {
	value string
}

type ShowDatabaseResponseDatabaseTypeEnum struct {
	MANAGED_DATABASE ShowDatabaseResponseDatabaseType
	FOREIGN_DATABASE ShowDatabaseResponseDatabaseType
}

func GetShowDatabaseResponseDatabaseTypeEnum() ShowDatabaseResponseDatabaseTypeEnum {
	return ShowDatabaseResponseDatabaseTypeEnum{
		MANAGED_DATABASE: ShowDatabaseResponseDatabaseType{
			value: "MANAGED_DATABASE",
		},
		FOREIGN_DATABASE: ShowDatabaseResponseDatabaseType{
			value: "FOREIGN_DATABASE",
		},
	}
}

func (c ShowDatabaseResponseDatabaseType) Value() string {
	return c.value
}

func (c ShowDatabaseResponseDatabaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatabaseResponseDatabaseType) UnmarshalJSON(b []byte) error {
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
