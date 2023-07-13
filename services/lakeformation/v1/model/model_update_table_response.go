package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateTableResponse Response Object
type UpdateTableResponse struct {

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
	OwnerType *UpdateTableResponseOwnerType `json:"owner_type,omitempty"`

	// 参数
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表分区列的列表
	PartitionKeys *[]Column `json:"partition_keys,omitempty"`

	// 表保留时间
	Retention *int32 `json:"retention,omitempty"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor,omitempty"`

	// 表类型
	TableType *UpdateTableResponseTableType `json:"table_type,omitempty"`

	// 表描述信息
	Comments *string `json:"comments,omitempty"`

	// 如果表是视图，则为视图的扩展文本；否则为 null
	ViewExpandedText *string `json:"view_expanded_text,omitempty"`

	// 如果表是视图，则为视图的原始文本；否则为 null
	ViewOriginalText *string `json:"view_original_text,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o UpdateTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateTableResponse", string(data)}, " ")
}

type UpdateTableResponseOwnerType struct {
	value string
}

type UpdateTableResponseOwnerTypeEnum struct {
	USER  UpdateTableResponseOwnerType
	GROUP UpdateTableResponseOwnerType
	ROLE  UpdateTableResponseOwnerType
}

func GetUpdateTableResponseOwnerTypeEnum() UpdateTableResponseOwnerTypeEnum {
	return UpdateTableResponseOwnerTypeEnum{
		USER: UpdateTableResponseOwnerType{
			value: "USER",
		},
		GROUP: UpdateTableResponseOwnerType{
			value: "GROUP",
		},
		ROLE: UpdateTableResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c UpdateTableResponseOwnerType) Value() string {
	return c.value
}

func (c UpdateTableResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTableResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type UpdateTableResponseTableType struct {
	value string
}

type UpdateTableResponseTableTypeEnum struct {
	MANAGED_TABLE     UpdateTableResponseTableType
	EXTERNAL_TABLE    UpdateTableResponseTableType
	VIRTUAL_VIEW      UpdateTableResponseTableType
	MATERIALIZED_VIEW UpdateTableResponseTableType
}

func GetUpdateTableResponseTableTypeEnum() UpdateTableResponseTableTypeEnum {
	return UpdateTableResponseTableTypeEnum{
		MANAGED_TABLE: UpdateTableResponseTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: UpdateTableResponseTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: UpdateTableResponseTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: UpdateTableResponseTableType{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c UpdateTableResponseTableType) Value() string {
	return c.value
}

func (c UpdateTableResponseTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTableResponseTableType) UnmarshalJSON(b []byte) error {
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
