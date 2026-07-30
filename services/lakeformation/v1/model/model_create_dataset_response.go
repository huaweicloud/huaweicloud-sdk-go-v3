package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateDatasetResponse Response Object
type CreateDatasetResponse struct {

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
	StorageType *CreateDatasetResponseStorageType `json:"storage_type,omitempty"`

	DatasetFormat *DatasetFileFormat `json:"dataset_format,omitempty"`

	// Dataset所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务一期实例响应Body无该参数。
	OwnerType *CreateDatasetResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务一期实例响应Body无该参数。
	OwnerSource *CreateDatasetResponseOwnerSource `json:"owner_source,omitempty"`

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

func (o CreateDatasetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasetResponse struct{}"
	}

	return strings.Join([]string{"CreateDatasetResponse", string(data)}, " ")
}

type CreateDatasetResponseStorageType struct {
	value string
}

type CreateDatasetResponseStorageTypeEnum struct {
	EXTERNAL CreateDatasetResponseStorageType
	MANAGED  CreateDatasetResponseStorageType
}

func GetCreateDatasetResponseStorageTypeEnum() CreateDatasetResponseStorageTypeEnum {
	return CreateDatasetResponseStorageTypeEnum{
		EXTERNAL: CreateDatasetResponseStorageType{
			value: "EXTERNAL",
		},
		MANAGED: CreateDatasetResponseStorageType{
			value: "MANAGED",
		},
	}
}

func (c CreateDatasetResponseStorageType) Value() string {
	return c.value
}

func (c CreateDatasetResponseStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatasetResponseStorageType) UnmarshalJSON(b []byte) error {
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

type CreateDatasetResponseOwnerType struct {
	value string
}

type CreateDatasetResponseOwnerTypeEnum struct {
	USER  CreateDatasetResponseOwnerType
	ROLE  CreateDatasetResponseOwnerType
	GROUP CreateDatasetResponseOwnerType
}

func GetCreateDatasetResponseOwnerTypeEnum() CreateDatasetResponseOwnerTypeEnum {
	return CreateDatasetResponseOwnerTypeEnum{
		USER: CreateDatasetResponseOwnerType{
			value: "USER",
		},
		ROLE: CreateDatasetResponseOwnerType{
			value: "ROLE",
		},
		GROUP: CreateDatasetResponseOwnerType{
			value: "GROUP",
		},
	}
}

func (c CreateDatasetResponseOwnerType) Value() string {
	return c.value
}

func (c CreateDatasetResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatasetResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type CreateDatasetResponseOwnerSource struct {
	value string
}

type CreateDatasetResponseOwnerSourceEnum struct {
	IAM         CreateDatasetResponseOwnerSource
	SAML        CreateDatasetResponseOwnerSource
	LDAP        CreateDatasetResponseOwnerSource
	LOCAL       CreateDatasetResponseOwnerSource
	AGENTTENANT CreateDatasetResponseOwnerSource
	OTHER       CreateDatasetResponseOwnerSource
}

func GetCreateDatasetResponseOwnerSourceEnum() CreateDatasetResponseOwnerSourceEnum {
	return CreateDatasetResponseOwnerSourceEnum{
		IAM: CreateDatasetResponseOwnerSource{
			value: "IAM",
		},
		SAML: CreateDatasetResponseOwnerSource{
			value: "SAML",
		},
		LDAP: CreateDatasetResponseOwnerSource{
			value: "LDAP",
		},
		LOCAL: CreateDatasetResponseOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: CreateDatasetResponseOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: CreateDatasetResponseOwnerSource{
			value: "OTHER",
		},
	}
}

func (c CreateDatasetResponseOwnerSource) Value() string {
	return c.value
}

func (c CreateDatasetResponseOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatasetResponseOwnerSource) UnmarshalJSON(b []byte) error {
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
