package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateDatasetResponse Response Object
type UpdateDatasetResponse struct {

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
	StorageType *UpdateDatasetResponseStorageType `json:"storage_type,omitempty"`

	DatasetFormat *DatasetFileFormat `json:"dataset_format,omitempty"`

	// Dataset所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务一期实例响应Body无该参数。
	OwnerType *UpdateDatasetResponseOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务一期实例响应Body无该参数。
	OwnerSource *UpdateDatasetResponseOwnerSource `json:"owner_source,omitempty"`

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

func (o UpdateDatasetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatasetResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatasetResponse", string(data)}, " ")
}

type UpdateDatasetResponseStorageType struct {
	value string
}

type UpdateDatasetResponseStorageTypeEnum struct {
	EXTERNAL UpdateDatasetResponseStorageType
	MANAGED  UpdateDatasetResponseStorageType
}

func GetUpdateDatasetResponseStorageTypeEnum() UpdateDatasetResponseStorageTypeEnum {
	return UpdateDatasetResponseStorageTypeEnum{
		EXTERNAL: UpdateDatasetResponseStorageType{
			value: "EXTERNAL",
		},
		MANAGED: UpdateDatasetResponseStorageType{
			value: "MANAGED",
		},
	}
}

func (c UpdateDatasetResponseStorageType) Value() string {
	return c.value
}

func (c UpdateDatasetResponseStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDatasetResponseStorageType) UnmarshalJSON(b []byte) error {
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

type UpdateDatasetResponseOwnerType struct {
	value string
}

type UpdateDatasetResponseOwnerTypeEnum struct {
	USER  UpdateDatasetResponseOwnerType
	ROLE  UpdateDatasetResponseOwnerType
	GROUP UpdateDatasetResponseOwnerType
}

func GetUpdateDatasetResponseOwnerTypeEnum() UpdateDatasetResponseOwnerTypeEnum {
	return UpdateDatasetResponseOwnerTypeEnum{
		USER: UpdateDatasetResponseOwnerType{
			value: "USER",
		},
		ROLE: UpdateDatasetResponseOwnerType{
			value: "ROLE",
		},
		GROUP: UpdateDatasetResponseOwnerType{
			value: "GROUP",
		},
	}
}

func (c UpdateDatasetResponseOwnerType) Value() string {
	return c.value
}

func (c UpdateDatasetResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDatasetResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type UpdateDatasetResponseOwnerSource struct {
	value string
}

type UpdateDatasetResponseOwnerSourceEnum struct {
	IAM         UpdateDatasetResponseOwnerSource
	SAML        UpdateDatasetResponseOwnerSource
	LDAP        UpdateDatasetResponseOwnerSource
	LOCAL       UpdateDatasetResponseOwnerSource
	AGENTTENANT UpdateDatasetResponseOwnerSource
	OTHER       UpdateDatasetResponseOwnerSource
}

func GetUpdateDatasetResponseOwnerSourceEnum() UpdateDatasetResponseOwnerSourceEnum {
	return UpdateDatasetResponseOwnerSourceEnum{
		IAM: UpdateDatasetResponseOwnerSource{
			value: "IAM",
		},
		SAML: UpdateDatasetResponseOwnerSource{
			value: "SAML",
		},
		LDAP: UpdateDatasetResponseOwnerSource{
			value: "LDAP",
		},
		LOCAL: UpdateDatasetResponseOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: UpdateDatasetResponseOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: UpdateDatasetResponseOwnerSource{
			value: "OTHER",
		},
	}
}

func (c UpdateDatasetResponseOwnerSource) Value() string {
	return c.value
}

func (c UpdateDatasetResponseOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDatasetResponseOwnerSource) UnmarshalJSON(b []byte) error {
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
