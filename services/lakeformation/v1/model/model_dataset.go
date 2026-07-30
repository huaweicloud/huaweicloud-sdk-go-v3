package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Dataset 文件集
type Dataset struct {

	// catalog名称
	CatalogName string `json:"catalog_name"`

	// catalogID
	CatalogId *string `json:"catalog_id,omitempty"`

	// 数据集名称
	DatasetName string `json:"dataset_name"`

	// DatasetID
	DatasetId *string `json:"dataset_id,omitempty"`

	// 数据集的描述信息
	Description *string `json:"description,omitempty"`

	// 数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库ID。
	DatabaseId *string `json:"database_id,omitempty"`

	// 数据集存储类型：EXTERNAL-外置存储,MANAGED-系统托管存储
	StorageType *DatasetStorageType `json:"storage_type,omitempty"`

	DatasetFormat *DatasetFileFormat `json:"dataset_format,omitempty"`

	// Dataset所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务一期实例响应Body无该参数。
	OwnerType *DatasetOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务一期实例响应Body无该参数。
	OwnerSource *DatasetOwnerSource `json:"owner_source,omitempty"`

	// 外置存储类型的元数据存储位置
	Location *string `json:"location,omitempty"`

	// 数据集其他属性
	Properties map[string]string `json:"properties,omitempty"`

	// 数据集创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 数据集修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o Dataset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dataset struct{}"
	}

	return strings.Join([]string{"Dataset", string(data)}, " ")
}

type DatasetStorageType struct {
	value string
}

type DatasetStorageTypeEnum struct {
	EXTERNAL DatasetStorageType
	MANAGED  DatasetStorageType
}

func GetDatasetStorageTypeEnum() DatasetStorageTypeEnum {
	return DatasetStorageTypeEnum{
		EXTERNAL: DatasetStorageType{
			value: "EXTERNAL",
		},
		MANAGED: DatasetStorageType{
			value: "MANAGED",
		},
	}
}

func (c DatasetStorageType) Value() string {
	return c.value
}

func (c DatasetStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasetStorageType) UnmarshalJSON(b []byte) error {
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

type DatasetOwnerType struct {
	value string
}

type DatasetOwnerTypeEnum struct {
	USER  DatasetOwnerType
	ROLE  DatasetOwnerType
	GROUP DatasetOwnerType
}

func GetDatasetOwnerTypeEnum() DatasetOwnerTypeEnum {
	return DatasetOwnerTypeEnum{
		USER: DatasetOwnerType{
			value: "USER",
		},
		ROLE: DatasetOwnerType{
			value: "ROLE",
		},
		GROUP: DatasetOwnerType{
			value: "GROUP",
		},
	}
}

func (c DatasetOwnerType) Value() string {
	return c.value
}

func (c DatasetOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasetOwnerType) UnmarshalJSON(b []byte) error {
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

type DatasetOwnerSource struct {
	value string
}

type DatasetOwnerSourceEnum struct {
	IAM         DatasetOwnerSource
	SAML        DatasetOwnerSource
	LDAP        DatasetOwnerSource
	LOCAL       DatasetOwnerSource
	AGENTTENANT DatasetOwnerSource
	OTHER       DatasetOwnerSource
}

func GetDatasetOwnerSourceEnum() DatasetOwnerSourceEnum {
	return DatasetOwnerSourceEnum{
		IAM: DatasetOwnerSource{
			value: "IAM",
		},
		SAML: DatasetOwnerSource{
			value: "SAML",
		},
		LDAP: DatasetOwnerSource{
			value: "LDAP",
		},
		LOCAL: DatasetOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: DatasetOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: DatasetOwnerSource{
			value: "OTHER",
		},
	}
}

func (c DatasetOwnerSource) Value() string {
	return c.value
}

func (c DatasetOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasetOwnerSource) UnmarshalJSON(b []byte) error {
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
