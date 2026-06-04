package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateTableInput 创建表输入模型
type CreateTableInput struct {

	// 表名称。只能包含中文、字母、数字和下划线，且长度为1~256个字符。
	TableName string `json:"table_name"`

	// 表格式，支持HIVE、ICEBERG、LANCE
	TableFormat *CreateTableInputTableFormat `json:"table_format,omitempty"`

	// 表类型，MANAGED_TABLE-内表、EXTERNAL_TABLE-外表、VIRTUAL_VIEW-视图、MATERIALIZED_VIEW-物化视图、DICTIONARY_TABLE-字典表、LAKE_TABLE-内表
	TableType CreateTableInputTableType `json:"table_type"`

	// 表所有者。只能包含字母、数字和下划线，且长度为0~49个字符。可以为null。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型，USER-用户、GROUP-组、ROLE-角色
	OwnerType CreateTableInputOwnerType `json:"owner_type"`

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

	// 表描述信息。由用户创建表时输入，最大长度为4000个字符。
	Comments *string `json:"comments,omitempty"`

	// 如果表是视图，则为视图的扩展文本；否则为 null
	ViewExpandedText *string `json:"view_expanded_text,omitempty"`

	// 如果表是视图，则为视图的原始文本；否则为 null
	ViewOriginalText *string `json:"view_original_text,omitempty"`

	// 是否忽略内表建表时对Obs路径的限制
	IgnoreObsChecked *bool `json:"ignore_obs_checked,omitempty"`

	// 数据概况统计开关。默认状态为开，修改table开关状态后，还需检查所属database的开关状态。例如：table与所属database开关同时打开，则数据概况统计开启。否则关闭
	DataStatisticEnable *bool `json:"data_statistic_enable,omitempty"`

	CreateOpenTableFormatInput *CreateOpenTableFormatInput `json:"create_open_table_format_input,omitempty"`
}

func (o CreateTableInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableInput struct{}"
	}

	return strings.Join([]string{"CreateTableInput", string(data)}, " ")
}

type CreateTableInputTableFormat struct {
	value string
}

type CreateTableInputTableFormatEnum struct {
	HIVE    CreateTableInputTableFormat
	ICEBERG CreateTableInputTableFormat
	LANCE   CreateTableInputTableFormat
}

func GetCreateTableInputTableFormatEnum() CreateTableInputTableFormatEnum {
	return CreateTableInputTableFormatEnum{
		HIVE: CreateTableInputTableFormat{
			value: "HIVE",
		},
		ICEBERG: CreateTableInputTableFormat{
			value: "ICEBERG",
		},
		LANCE: CreateTableInputTableFormat{
			value: "LANCE",
		},
	}
}

func (c CreateTableInputTableFormat) Value() string {
	return c.value
}

func (c CreateTableInputTableFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableInputTableFormat) UnmarshalJSON(b []byte) error {
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

type CreateTableInputTableType struct {
	value string
}

type CreateTableInputTableTypeEnum struct {
	MANAGED_TABLE     CreateTableInputTableType
	EXTERNAL_TABLE    CreateTableInputTableType
	VIRTUAL_VIEW      CreateTableInputTableType
	MATERIALIZED_VIEW CreateTableInputTableType
	DICTIONARY_TABLE  CreateTableInputTableType
	LAKE_TABLE        CreateTableInputTableType
}

func GetCreateTableInputTableTypeEnum() CreateTableInputTableTypeEnum {
	return CreateTableInputTableTypeEnum{
		MANAGED_TABLE: CreateTableInputTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: CreateTableInputTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: CreateTableInputTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: CreateTableInputTableType{
			value: "MATERIALIZED_VIEW",
		},
		DICTIONARY_TABLE: CreateTableInputTableType{
			value: "DICTIONARY_TABLE",
		},
		LAKE_TABLE: CreateTableInputTableType{
			value: "LAKE_TABLE",
		},
	}
}

func (c CreateTableInputTableType) Value() string {
	return c.value
}

func (c CreateTableInputTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableInputTableType) UnmarshalJSON(b []byte) error {
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

type CreateTableInputOwnerType struct {
	value string
}

type CreateTableInputOwnerTypeEnum struct {
	USER  CreateTableInputOwnerType
	GROUP CreateTableInputOwnerType
	ROLE  CreateTableInputOwnerType
}

func GetCreateTableInputOwnerTypeEnum() CreateTableInputOwnerTypeEnum {
	return CreateTableInputOwnerTypeEnum{
		USER: CreateTableInputOwnerType{
			value: "USER",
		},
		GROUP: CreateTableInputOwnerType{
			value: "GROUP",
		},
		ROLE: CreateTableInputOwnerType{
			value: "ROLE",
		},
	}
}

func (c CreateTableInputOwnerType) Value() string {
	return c.value
}

func (c CreateTableInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableInputOwnerType) UnmarshalJSON(b []byte) error {
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
