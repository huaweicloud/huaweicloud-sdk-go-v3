package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// TableInput 表输入模型
type TableInput struct {

	// 表名字
	TableName string `json:"table_name"`

	// 表类型
	TableType TableInputTableType `json:"table_type"`

	// 表所有者
	Owner string `json:"owner"`

	// 所有者类型
	OwnerType TableInputOwnerType `json:"owner_type"`

	// 表创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 最近一次访问时间
	LastAccessTime *sdktime.SdkTime `json:"last_access_time,omitempty"`

	// 最近一次分析统计时间
	LastAnalyzedTime *sdktime.SdkTime `json:"last_analyzed_time,omitempty"`

	// 分区列的信息
	PartitionKeys *[]Column `json:"partition_keys,omitempty"`

	// 表保留时间
	Retention *int32 `json:"retention,omitempty"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor"`

	// 表参数信息，每个键是一个键字符串，不少于 1 个字节或超过 255 个字节 每个值是一个 UTF-8 字符串，不超过 4000 个字节
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表描述信息
	Comments *string `json:"comments,omitempty"`

	// 如果表是视图，则为视图的扩展文本；否则为 null
	ViewExpandedText *string `json:"view_expanded_text,omitempty"`

	// 如果表是视图，则为视图的原始文本；否则为 null
	ViewOriginalText *string `json:"view_original_text,omitempty"`
}

func (o TableInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableInput struct{}"
	}

	return strings.Join([]string{"TableInput", string(data)}, " ")
}

type TableInputTableType struct {
	value string
}

type TableInputTableTypeEnum struct {
	MANAGED_TABLE     TableInputTableType
	EXTERNAL_TABLE    TableInputTableType
	VIRTUAL_VIEW      TableInputTableType
	MATERIALIZED_VIEW TableInputTableType
}

func GetTableInputTableTypeEnum() TableInputTableTypeEnum {
	return TableInputTableTypeEnum{
		MANAGED_TABLE: TableInputTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: TableInputTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: TableInputTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: TableInputTableType{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c TableInputTableType) Value() string {
	return c.value
}

func (c TableInputTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableInputTableType) UnmarshalJSON(b []byte) error {
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

type TableInputOwnerType struct {
	value string
}

type TableInputOwnerTypeEnum struct {
	USER  TableInputOwnerType
	GROUP TableInputOwnerType
	ROLE  TableInputOwnerType
}

func GetTableInputOwnerTypeEnum() TableInputOwnerTypeEnum {
	return TableInputOwnerTypeEnum{
		USER: TableInputOwnerType{
			value: "USER",
		},
		GROUP: TableInputOwnerType{
			value: "GROUP",
		},
		ROLE: TableInputOwnerType{
			value: "ROLE",
		},
	}
}

func (c TableInputOwnerType) Value() string {
	return c.value
}

func (c TableInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableInputOwnerType) UnmarshalJSON(b []byte) error {
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
