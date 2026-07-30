package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowDatasetResponse Response Object
type ShowDatasetResponse struct {

	// catalog名称
	CatalogName *string `json:"catalog_name,omitempty"`

	// catalogID
	CatalogId *string `json:"catalog_id,omitempty"`

	// 数据集名称
	DatasetName *string `json:"dataset_name,omitempty"`

	// DatasetID
	DatasetId *string `json:"dataset_id,omitempty"`

	// 数据集的描述信息
	Description *string `json:"description,omitempty"`

	// 数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库ID。
	DatabaseId *string `json:"database_id,omitempty"`

	// 数据集存储类型：EXTERNAL-外置存储,MANAGED-系统托管存储
	StorageType *ShowDatasetResponseStorageType `json:"storage_type,omitempty"`

	DatasetFormat *DatasetFileFormat `json:"dataset_format,omitempty"`

	// Dataset所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务一期实例响应Body无该参数。
	OwnerType *ShowDatasetResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务一期实例响应Body无该参数。
	OwnerSource *ShowDatasetResponseOwnerSource `json:"owner_source,omitempty"`

	// 外置存储类型的元数据存储位置
	Location *string `json:"location,omitempty"`

	// 数据集其他属性
	Properties map[string]string `json:"properties,omitempty"`

	// 数据集创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 数据集修改时间
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDatasetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatasetResponse struct{}"
	}

	return strings.Join([]string{"ShowDatasetResponse", string(data)}, " ")
}

type ShowDatasetResponseStorageType struct {
	value string
}

type ShowDatasetResponseStorageTypeEnum struct {
	EXTERNAL ShowDatasetResponseStorageType
	MANAGED  ShowDatasetResponseStorageType
}

func GetShowDatasetResponseStorageTypeEnum() ShowDatasetResponseStorageTypeEnum {
	return ShowDatasetResponseStorageTypeEnum{
		EXTERNAL: ShowDatasetResponseStorageType{
			value: "EXTERNAL",
		},
		MANAGED: ShowDatasetResponseStorageType{
			value: "MANAGED",
		},
	}
}

func (c ShowDatasetResponseStorageType) Value() string {
	return c.value
}

func (c ShowDatasetResponseStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatasetResponseStorageType) UnmarshalJSON(b []byte) error {
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

type ShowDatasetResponseOwnerType struct {
	value string
}

type ShowDatasetResponseOwnerTypeEnum struct {
	USER  ShowDatasetResponseOwnerType
	ROLE  ShowDatasetResponseOwnerType
	GROUP ShowDatasetResponseOwnerType
}

func GetShowDatasetResponseOwnerTypeEnum() ShowDatasetResponseOwnerTypeEnum {
	return ShowDatasetResponseOwnerTypeEnum{
		USER: ShowDatasetResponseOwnerType{
			value: "USER",
		},
		ROLE: ShowDatasetResponseOwnerType{
			value: "ROLE",
		},
		GROUP: ShowDatasetResponseOwnerType{
			value: "GROUP",
		},
	}
}

func (c ShowDatasetResponseOwnerType) Value() string {
	return c.value
}

func (c ShowDatasetResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatasetResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type ShowDatasetResponseOwnerSource struct {
	value string
}

type ShowDatasetResponseOwnerSourceEnum struct {
	IAM         ShowDatasetResponseOwnerSource
	SAML        ShowDatasetResponseOwnerSource
	LDAP        ShowDatasetResponseOwnerSource
	LOCAL       ShowDatasetResponseOwnerSource
	AGENTTENANT ShowDatasetResponseOwnerSource
	OTHER       ShowDatasetResponseOwnerSource
}

func GetShowDatasetResponseOwnerSourceEnum() ShowDatasetResponseOwnerSourceEnum {
	return ShowDatasetResponseOwnerSourceEnum{
		IAM: ShowDatasetResponseOwnerSource{
			value: "IAM",
		},
		SAML: ShowDatasetResponseOwnerSource{
			value: "SAML",
		},
		LDAP: ShowDatasetResponseOwnerSource{
			value: "LDAP",
		},
		LOCAL: ShowDatasetResponseOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: ShowDatasetResponseOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: ShowDatasetResponseOwnerSource{
			value: "OTHER",
		},
	}
}

func (c ShowDatasetResponseOwnerSource) Value() string {
	return c.value
}

func (c ShowDatasetResponseOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDatasetResponseOwnerSource) UnmarshalJSON(b []byte) error {
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
