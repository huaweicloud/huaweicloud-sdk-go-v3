package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Table table
type Table struct {

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名字
	TableName string `json:"table_name"`

	// 表创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 最后访问时间
	LastAccessTime *sdktime.SdkTime `json:"last_access_time"`

	// 表元数据最后一次修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 上一次做列级别的统计信息计算的时间
	LastAnalyzedTime *sdktime.SdkTime `json:"last_analyzed_time,omitempty"`

	// 用户信息
	Owner string `json:"owner"`

	// 用户类型
	OwnerType TableOwnerType `json:"owner_type"`

	// 参数
	Parameters map[string]string `json:"parameters"`

	// 表分区列的列表
	PartitionKeys []Column `json:"partition_keys"`

	// 表保留时间
	Retention int32 `json:"retention"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor"`

	// 表类型
	TableType TableTableType `json:"table_type"`

	// 表描述信息
	Comments *string `json:"comments,omitempty"`

	// 如果表是视图，则为视图的扩展文本；否则为 null
	ViewExpandedText string `json:"view_expanded_text"`

	// 如果表是视图，则为视图的原始文本；否则为 null
	ViewOriginalText string `json:"view_original_text"`
}

func (o Table) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Table struct{}"
	}

	return strings.Join([]string{"Table", string(data)}, " ")
}

type TableOwnerType struct {
	value string
}

type TableOwnerTypeEnum struct {
	USER  TableOwnerType
	GROUP TableOwnerType
	ROLE  TableOwnerType
}

func GetTableOwnerTypeEnum() TableOwnerTypeEnum {
	return TableOwnerTypeEnum{
		USER: TableOwnerType{
			value: "USER",
		},
		GROUP: TableOwnerType{
			value: "GROUP",
		},
		ROLE: TableOwnerType{
			value: "ROLE",
		},
	}
}

func (c TableOwnerType) Value() string {
	return c.value
}

func (c TableOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableOwnerType) UnmarshalJSON(b []byte) error {
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

type TableTableType struct {
	value string
}

type TableTableTypeEnum struct {
	MANAGED_TABLE     TableTableType
	EXTERNAL_TABLE    TableTableType
	VIRTUAL_VIEW      TableTableType
	MATERIALIZED_VIEW TableTableType
}

func GetTableTableTypeEnum() TableTableTypeEnum {
	return TableTableTypeEnum{
		MANAGED_TABLE: TableTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: TableTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: TableTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: TableTableType{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c TableTableType) Value() string {
	return c.value
}

func (c TableTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableTableType) UnmarshalJSON(b []byte) error {
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
