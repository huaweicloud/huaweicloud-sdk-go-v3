package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type CreateTableResponse struct {

	// catalog名字
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名字
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名字
	TableName *string `json:"table_name,omitempty"`

	// 表创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 最后访问时间
	LastAccessTime *sdktime.SdkTime `json:"last_access_time,omitempty"`

	// 表元数据最后一次修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 上一次做列级别的统计信息计算的时间
	LastAnalyzedTime *sdktime.SdkTime `json:"last_analyzed_time,omitempty"`

	// 用户信息
	Owner *string `json:"owner,omitempty"`

	// 用户类型
	OwnerType *CreateTableResponseOwnerType `json:"owner_type,omitempty"`

	// 参数
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表分区列的列表
	PartitionKeys *[]Column `json:"partition_keys,omitempty"`

	// 表保留时间
	Retention *int32 `json:"retention,omitempty"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor,omitempty"`

	// 表类型
	TableType *CreateTableResponseTableType `json:"table_type,omitempty"`

	// 表描述信息
	Comments *string `json:"comments,omitempty"`

	// 如果表是视图，则为视图的扩展文本；否则为 null
	ViewExpandedText *string `json:"view_expanded_text,omitempty"`

	// 如果表是视图，则为视图的原始文本；否则为 null
	ViewOriginalText *string `json:"view_original_text,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o CreateTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableResponse struct{}"
	}

	return strings.Join([]string{"CreateTableResponse", string(data)}, " ")
}

type CreateTableResponseOwnerType struct {
	value string
}

type CreateTableResponseOwnerTypeEnum struct {
	USER  CreateTableResponseOwnerType
	GROUP CreateTableResponseOwnerType
	ROLE  CreateTableResponseOwnerType
}

func GetCreateTableResponseOwnerTypeEnum() CreateTableResponseOwnerTypeEnum {
	return CreateTableResponseOwnerTypeEnum{
		USER: CreateTableResponseOwnerType{
			value: "USER",
		},
		GROUP: CreateTableResponseOwnerType{
			value: "GROUP",
		},
		ROLE: CreateTableResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c CreateTableResponseOwnerType) Value() string {
	return c.value
}

func (c CreateTableResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableResponseOwnerType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateTableResponseTableType struct {
	value string
}

type CreateTableResponseTableTypeEnum struct {
	MANAGED_TABLE     CreateTableResponseTableType
	EXTERNAL_TABLE    CreateTableResponseTableType
	VIRTUAL_VIEW      CreateTableResponseTableType
	MATERIALIZED_VIEW CreateTableResponseTableType
}

func GetCreateTableResponseTableTypeEnum() CreateTableResponseTableTypeEnum {
	return CreateTableResponseTableTypeEnum{
		MANAGED_TABLE: CreateTableResponseTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: CreateTableResponseTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: CreateTableResponseTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: CreateTableResponseTableType{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c CreateTableResponseTableType) Value() string {
	return c.value
}

func (c CreateTableResponseTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableResponseTableType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
